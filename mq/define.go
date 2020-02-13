package mq

// TransferData : 将要写到rabbitmq的数据的结构体(定义消息载体的结构)
type TransferData struct {
	FileHash      string
	CurLocation   string
	DestLocation  string
	//other
}
