package message

type Message struct {
	//发送时间 纳秒
	Ns int64
	//发送内容
	Msg string
	//发送者标识
	Publisher string
	//消息标识
	Mid string
}
