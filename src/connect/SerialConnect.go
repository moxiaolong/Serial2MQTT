package connect

import (
	"github.com/goburrow/serial"
	Config "modbusRtu2Mqtt/src/config"
	log "modbusRtu2Mqtt/src/userlog"
	"time"
)

func Serial(config Config.Config, serialChan chan serial.Port, serialHandle func(chan []byte)) {
	for {
		log.Info("Connecting Serial ", config.Address)
		port, err := serial.Open(&serial.Config{
			Address:  config.Address,
			BaudRate: config.BaudRate,
			DataBits: config.DataBits,
			StopBits: config.StopBits,
			Parity:   config.Parity,
			Timeout:  time.Second * 5,
		})
		if err != nil {
			log.Error("Connect Serial Error ", err)
			continue
		} else {
			log.Info("Connected Serial to ", config.Address)
			serialChan <- port
			bytes := make(chan []byte, 5)
			serialHandle(bytes)
			for {
				buffer := make([]byte, 256)
				read, err := port.Read(buffer)
				if err != nil {
					log.Fatal(err)
				} else {
					if read > 0 {
						log.Debug(buffer)
						bytes <- buffer
					}
				}
			}
		}

	}
}
