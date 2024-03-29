package connect

import (
	"encoding/hex"
	"github.com/goburrow/serial"
	"log"
	Config "modbusRtu2Mqtt/src/config"
	"time"
)

func Serial(config Config.Config, serialChan chan serial.Port, serialHandle func([]byte, Config.Config)) {
	for {
		log.Println("Connecting... Serial ", config.Address)
		port, err := serial.Open(&serial.Config{
			Address:  config.Address,
			BaudRate: config.BaudRate,
			DataBits: config.DataBits,
			StopBits: config.StopBits,
			Parity:   config.Parity,
			Timeout:  time.Second * 15,
		})
		if err != nil {
			log.Println("Connect Serial Error ", err)
			time.Sleep(time.Second * 15)
			continue
		} else {
			log.Println("Connected Serial to ", config.Address)
			serialChan <- port
			buffer := make([]byte, config.Serial.BufferSize)
			for {
				read, err := port.Read(buffer)

				if err != nil {
					log.Println(err)
				} else {
					if read > 0 {
						slice := buffer[0:read]
						log.Println("originalData :", hex.EncodeToString(slice))
						buffer = make([]byte, config.Serial.BufferSize)
						serialHandle(slice, config)
					}
				}
			}
		}
	}
}
