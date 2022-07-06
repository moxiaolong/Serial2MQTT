package handle

import (
	"fmt"
	"log"
)

func Serial(msg []byte) {
	log.Println("receive serial data:", fmt.Sprintf("%x", msg))

}
