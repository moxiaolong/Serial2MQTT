package connect

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	Config "modbusRtu2Mqtt/src/config"
	"strconv"
	"time"
)

func Mqtt(config Config.Config, clientId string, dealMqttMsg func(chan [2]string, chan bool), clients chan MQTT.Client) {

	log.Println("MQTT ClientId", clientId)
	username := config.Mqtt.UserName
	password := config.Mqtt.PassWord
	qos := config.Mqtt.Qos
	DownTopic := config.Mqtt.DownTopic
	server := "tcp://" + config.Mqtt.Host + ":" + strconv.Itoa(config.Mqtt.Port)
	connOpts := MQTT.NewClientOptions().AddBroker(server).SetClientID(clientId).SetCleanSession(true)
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

	MqttClient := MQTT.NewClient(connOpts)

	for {
		if MqttClient.IsConnected() {
			time.Sleep(time.Second * 15)
			continue
		}
		log.Println("Connecting... to mqtt:", server)
		if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			time.Sleep(time.Second * 15)
			continue
		} else {
			log.Println("Connected to mqtt:", server)
			clients <- MqttClient
			if token := MqttClient.Subscribe(DownTopic, byte(qos), nil); token.Wait() && token.Error() != nil {
				log.Println(token.Error())
			} else {
				log.Println("Subscribe mqtt topic successful :", DownTopic)
			}

		}
	}

}
