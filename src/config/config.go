package config

import (
	"github.com/spf13/viper"
	log "modbusRtu2Mqtt/src/userlog"
)

type Config struct {
	Mqtt
	Serial
}

type Mqtt struct {
	Host      string
	Port      int
	UserName  string
	PassWord  string
	Qos       int
	UpTopic   string
	DownTopic string
}

type Serial struct {
	//串口
	Address  string
	BaudRate int
	DataBits int
	StopBits int
	Parity   string
}

func GetConfig() Config {
	config := Config{}
	config.Mqtt.Host = "127.0.0.1"
	config.Mqtt.Port = 1883
	config.Mqtt.Qos = 0
	config.Mqtt.UpTopic = "/twwg/modbus_rtu/up"
	config.Mqtt.DownTopic = "/twwg/modbus_rtu/down"

	//modbus配置
	config.Serial.Address = "/dev/ttyUSB0"
	config.Serial.BaudRate = 115200
	config.Serial.DataBits = 8
	config.Serial.StopBits = 1
	config.Serial.Parity = "N"

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {

		log.Warn(err.Error())

		return config
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Warn("unable to decode into struct, %v", err)
	}

	return config
}
