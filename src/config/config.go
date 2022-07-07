package config

import (
	"github.com/spf13/viper"
	"log"
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
	Address          string
	BaudRate         int
	DataBits         int
	StopBits         int
	Parity           string
	BufferSize       int
	UnpackBufferSize int
	//包头
	UnpackBHead []byte
}

type Store interface {
	GetConfig()
}

var config Config = Config{}

func GetConfig() Config {
	config = Config{}
	config.Mqtt.Host = "127.0.0.1"
	config.Mqtt.Port = 1883
	config.Mqtt.Qos = 0
	config.Mqtt.UpTopic = "/twwg/serial/up"
	config.Mqtt.DownTopic = "/twwg/serial/down"

	//modbus配置
	config.Serial.Address = "/dev/ttyUSB0"
	config.Serial.BaudRate = 115200
	config.Serial.DataBits = 8
	config.Serial.StopBits = 1
	config.Serial.Parity = "N"
	config.Serial.BufferSize = 16
	config.Serial.UnpackBufferSize = 512
	config.Serial.UnpackBHead = []byte{0x24, 0x50, 0x41, 0x52, 0x41}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {

		log.Println(err.Error())

		return config
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Println("unable to decode into struct, %v", err)
	}

	return config
}
