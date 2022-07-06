package handle

import (
	"encoding/json"
	"modbusRtu2Mqtt/src/message"
	log "modbusRtu2Mqtt/src/userlog"
	"time"
)

func Mqtt(msg chan [2]string, exit chan bool) {
	for {
		select {
		case incoming := <-msg:
			log.Info("Received message on topic: %s\nMessage: %s\n", incoming[0], incoming[1])
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
