package handle

import (
	"encoding/hex"
	"encoding/json"
	"github.com/goburrow/serial"
	"log"
	"modbusRtu2Mqtt/src/message"
	"time"
)

var Port serial.Port

func SetSerial(port serial.Port) {
	Port = port
}

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
			log.Println("mqtt msg:", msg)
			decodeString, err := hex.DecodeString(m)
			if err != nil {
				log.Println("MQTT hex decode Error:", err)
			}
			writeMsg, err := Port.Write(decodeString)
			if err != nil {
				log.Println("Serial Write Error:", err)
			} else {
				log.Println("Serial Write Done:", writeMsg)
			}

		case <-exit:
			return
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}
