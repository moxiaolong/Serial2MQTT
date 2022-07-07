package handle

import (
	"encoding/hex"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	Config "modbusRtu2Mqtt/src/config"
	"modbusRtu2Mqtt/src/message"
	"time"
)

var buffer = make([]byte, 0)
var unpackChan = make(chan []byte)
var Client MQTT.Client

func SetMqttClient(client MQTT.Client) {
	Client = client
}

func Serial(msg []byte, config Config.Config) {

	buffer = append(buffer, msg...)
	go consumerUnpack(config)

	//拆包
	//包头
	head := config.UnpackBHead
	//在buffer中寻找head

	for {
		if len(buffer) < len(head) {
			//不足head长度
			return
		}
		//从头通过滑动窗口寻找
		oneMore := false
		//滑动窗口寻找头
		//从头后开始 到 buffer末尾和head末尾对齐结束
		for i := len(head) - 1; i < len(buffer)-len(head); i++ {
			//包含头
			isHead := true
			//比较每一位
			//遍历head
			for j := 0; j < len(head); j++ {
				if buffer[i+j] != head[j] {
					//不相等
					isHead = false
					break
				}
			}
			//是头
			if isHead {
				//完全相等
				//将head前的视作一个包 截取
				endIndex := i
				unpackChan <- append(buffer[:endIndex])
				buffer = append(buffer[endIndex:])
				//标记oneMore 之后跳出
				oneMore = true
				break
			}
			//继续位移窗口
		}
		if !oneMore {
			time.Sleep(time.Second * 1)
			break
		}
	}
}

func consumerUnpack(config Config.Config) {
	for bytes := range unpackChan {
		sprintf := hex.EncodeToString(bytes)
		log.Println("Unpacked Data:", sprintf)
		m := message.Message{Ns: time.Now().UnixNano(), Msg: sprintf}
		Client.Publish(config.Mqtt.UpTopic, 1, false, m)
		log.Println("Mqtt Publish Done :", m)
	}
}
