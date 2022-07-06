package handle

import (
	"encoding/json"
	"log"
	"modbusRtu2Mqtt/src/message"
	"time"
)

func Mqtt(msg chan [2]string, exit chan bool) {
	for {
		select {
		case incoming := <-msg:
			log.Println("Received message on topic: ", incoming[0], "Message:", incoming[1])
			var msg message.Message
			err := json.Unmarshal([]byte(incoming[1]), &msg)
			if err != nil {
				log.Println(err)
				continue
			}
			m := msg.Msg
			log.Println("mqtt msg:", m)
		case <-exit:
			return
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}
