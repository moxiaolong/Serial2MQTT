package connect

import (
	"github.com/goburrow/serial"
	"log"
	Config "modbusRtu2Mqtt/src/config"
	"time"
)

func Serial(config Config.Config, serialChan chan serial.Port, serialHandle func([]byte)) {
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
			buffer := make([]byte, 8)
			for {
				read, err := port.Read(buffer)
				if err != nil {
					log.Println(err)
				} else {
					if read > 0 {
						serialHandle(buffer)
					}
				}
			}
		}

	}
}
