package handle

import (
	"fmt"
	log "modbusRtu2Mqtt/src/userlog"
)

func Serial(msg chan []byte) {
	for bytes := range msg {
		sprintf := fmt.Sprintf("%x", bytes)
		log.Info("收到串口数据：", sprintf)
	}
}
