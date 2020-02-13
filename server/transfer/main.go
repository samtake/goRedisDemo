package main

import (
	"encoding/json"
	"goRedisDemo/config"
	"goRedisDemo/mq"
	"log"
	"os"
)

// ProcessTransfer : 处理文件转移
func ProcessTransfer(msg []byte) bool {
	log.Println(string(msg))

	//1.解析msg
	pubData := mq.TransferData{}
	err := json.Unmarshal(msg, &pubData)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	// 2.根据临时存储文件路径，创建文件句柄
	_, err = os.Open(pubData.CurLocation)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	// // 3.通过文件句柄将文件内容读出来并且上传到OSS
	// err = oss.Bucket().PutObject(
	// 	pubData.DestLocation,
	// 	bufio.NewReader(fin))
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return false
	// }

	// 4.更新文件的存储路径到文件表
	// _ = dblayer.UpdateFileLocation(
	// 	pubData.FileHash,
	// 	pubData.DestLocation)
	return true
}

func main() {
	if !config.AsyncTransferEnable {
		log.Println("异步转移文件功能目前被禁用，请检查相关配置")
		return
	}
	log.Println("文件转移服务启动中，开始监听转移任务队列...")
	mq.StartConsume(
		config.TransOSSQueueName,
		"transfer_oss",
		ProcessTransfer)
}
