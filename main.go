package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/goburrow/serial"
	"github.com/things-go/go-modbus"
	"math"
	"math/rand"
	Config "modbusRtu2Mqtt/src/config"
	"modbusRtu2Mqtt/src/message"
	log "modbusRtu2Mqtt/src/userlog"
	"strconv"
	"time"
)

var ClientID = "ModBusConverter" + strconv.FormatInt(rand.Int63n(math.MaxInt64), 16)
var MqttClient MQTT.Client = nil
var ModbusRtuClient modbus.Client = nil

// 处理订阅到的MQTT消息
func dealMqttMsg(msg chan [2]string, exit chan bool) {
	for {
		select {
		case incoming := <-msg:
			fmt.Printf("Received message on topic: %s\nMessage: %s\n", incoming[0], incoming[1])
			var msg message.Message
			err := json.Unmarshal([]byte(incoming[1]), &msg)
			if err != nil {
				log.Error(err)
				continue
			}
			m := msg.Msg
			log.Info(m)

		case <-exit:
			return
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func main() {
	config := Config.GetConfig()
	log.Info("当前配置", config)
	go connectMqtt(config)
	go connectModBusRtu(config)
	for {
		time.Sleep(time.Second * 60)
	}

}

func connectModBusRtu(config Config.Config) {
	address := config.RtuModbus.Address
	rate := config.RtuModbus.BaudRate
	bits := config.RtuModbus.DataBits
	stopBits := config.RtuModbus.StopBits
	parity := config.RtuModbus.Parity
	p := modbus.NewRTUClientProvider(modbus.WithEnableLogger(),
		modbus.WithSerialConfig(serial.Config{
			Address:  address,
			BaudRate: rate,
			DataBits: bits,
			StopBits: stopBits,
			Parity:   parity,
			Timeout:  modbus.SerialDefaultTimeout,
		}))

	ModbusRtuClient = modbus.NewClient(p)
	defer ModbusRtuClient.Close()
	for {
		err := ModbusRtuClient.Connect()
		if err != nil {
			log.Error("modbus rtu connect failed, ", err)
			time.Sleep(time.Second * 5)
			continue
		}
		log.Info("modbus starting")
		for {
			results, err := ModbusRtuClient.ReadCoils(3, 0, 10)
			if err != nil {
				log.Error(err.Error())
			}

			log.Info("ReadDiscreteInputs", results)

			time.Sleep(time.Second * 2)
		}
	}

}

func connectMqtt(config Config.Config) {
	log.Info("当前ClientId", ClientID)
	username := config.Mqtt.UserName
	password := config.Mqtt.PassWord
	qos := config.Mqtt.Qos
	DownTopic := config.Mqtt.DownTopic
	server := "tcp://" + config.Mqtt.Host + ":" + strconv.Itoa(config.Mqtt.Port)
	connOpts := MQTT.NewClientOptions().AddBroker(server).SetClientID(ClientID).SetCleanSession(true)
	if username != "" {
		connOpts.SetUsername(username)
		if password != "" {
			connOpts.SetPassword(password)
		}
	}
	connOpts.SetAutoReconnect(true)
	connOpts.SetMaxReconnectInterval(5)

	quit := make(chan bool)
	recmsg := make(chan [2]string, 300)

	connOpts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		recmsg <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	go dealMqttMsg(recmsg, quit)

	MqttClient = MQTT.NewClient(connOpts)

	for {
		if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
			log.Error(token.Error())
			time.Sleep(time.Second * 5)
			continue
		} else {
			log.Info("Connected to :", server)
			if token := MqttClient.Subscribe(DownTopic, byte(qos), nil); token.Wait() && token.Error() != nil {
				log.Error(token.Error())
			} else {
				log.Info("Subscribe topic successful :", DownTopic)
			}
			return
		}

	}

}
