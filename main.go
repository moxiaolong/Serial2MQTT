package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/goburrow/serial"
	"math"
	"math/rand"
	Config "modbusRtu2Mqtt/src/config"
	"modbusRtu2Mqtt/src/connect"
	"modbusRtu2Mqtt/src/handle"
	log "modbusRtu2Mqtt/src/userlog"
	"strconv"
	"time"
)

var ClientID = "SerialMqttConverter" + strconv.FormatInt(rand.Int63n(math.MaxInt64), 16)
var MqttClient MQTT.Client = nil
var serialPort serial.Port = nil

func main() {

	config := Config.GetConfig()
	log.Info("config : ", config)
	clients := make(chan MQTT.Client, 1)
	go connect.Mqtt(config, ClientID, handle.Mqtt, clients)
	ports := make(chan serial.Port, 1)
	go connect.Serial(config, ports, handle.Serial)

	MqttClient = <-clients
	serialPort = <-ports
	for {
		time.Sleep(time.Second * 60)
	}
}
