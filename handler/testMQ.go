package handler

import (
	"encoding/json"
	mq "goRedisDemo/mq"
)

// UploadHandler ：
func UploadHandlerMQ() {
	// 写入异步转移任务队列
	data := mq.TransferData{
		FileHash:     "fileMeta.FileSha1",
		CurLocation:  "fileMeta.Location",
		DestLocation: "ossPath",
		// DestStoreType: cmn.StoreOSS,
	}
	pubData, _ := json.Marshal(data)
	pubSuc := mq.Publish(
		"cfg.TransExchangeName",
		"cfg.TransOSSRoutingKey",
		pubData,
	)
	if !pubSuc {
		// TODO: 当前发送转移信息失败，稍后重试
	}
}
