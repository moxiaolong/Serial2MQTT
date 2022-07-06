package handle

import (
	"fmt"
	"log"
)

func Serial(msg []byte) {

	for _, bytes := range msg {
		sprintf := fmt.Sprintf("%x", bytes)
		log.Println("receive serial data:：", sprintf)
	}
}
