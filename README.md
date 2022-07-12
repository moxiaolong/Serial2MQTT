# Serial2MQTT

- 接受串口数据，拆包后发送到MQTT
- 接受MQTT数据，发送到串口

通过config.yaml配置

```yaml
Mqtt:
  Host: 127.0.0.1
  Port: 1883
  Qos: 0
  UpTopic: /twwg/serial/up #上传topic
  DownTopic: /twwg/serial/down #下发topic
Serial:
  Address: /dev/ttyUSB0 #串口地址
  BaudRate: 115200
  DataBits: 8
  StopBits: 1
  Parity: N
  BufferSize: 16 #串口每次读取长度
  UnpackBufferSize: 512 #拆包缓冲区长度
  UnpackBHead: 0x24, 0x50, 0x41, 0x52, 0x41 #包头，用来拆包
```

上下行数据格式：
{"Ns":1657188619495246638,"Msg":"ffffffff"}
