package connect

import (
	"github.com/goburrow/serial"
	Config "modbusRtu2Mqtt/src/config"
	log "modbusRtu2Mqtt/src/userlog"
	"time"
)

func Serial(config Config.Config, serialChan chan serial.Port, serialHandle func(*[]byte)) {
	for {
		log.Info("Connecting... Serial ", config.Address)
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
			buffer := make([]byte, 4)
			serialHandle(&buffer)
			for {
				read, err := port.Read(buffer)
				if err != nil {
					log.Warn(err)
				} else {
					if read == 0 {
						time.Sleep(time.Second * 1)
					}
				}
			}
		}

	}
}
