package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/goburrow/serial"
	"log"
	"math"
	"math/rand"
	Config "modbusRtu2Mqtt/src/config"
	"modbusRtu2Mqtt/src/connect"
	"modbusRtu2Mqtt/src/handle"
	"strconv"
	"time"
)

func main() {
	//日志
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	config := Config.GetConfig()

	log.Println("config : ", config)

	//连接mqtt
	var ClientID = "SerialMqttConverter" + strconv.FormatInt(rand.Int63n(math.MaxInt64), 16)
	clients := make(chan MQTT.Client, 1)
	go connect.Mqtt(config, ClientID, handle.Mqtt, clients)

	//连接串口
	ports := make(chan serial.Port, 1)
	go connect.Serial(config, ports, handle.Serial)

	MqttClient := <-clients
	SerialPort := <-ports

	handle.SetSerial(SerialPort)
	handle.SetMqttClient(MqttClient)

	for {
		time.Sleep(time.Second * 60)
	}
}
