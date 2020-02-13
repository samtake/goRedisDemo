package main

import (
	"fmt"
	cfg "goRedisDemo/config"
	"net/http"
)

func main() {
	fmt.Printf("上传服务启动中，开始监听监听[%s]...\n", cfg.UploadServiceHost)
	// 启动服务并监听端口
	err := http.ListenAndServe(cfg.UploadServiceHost, nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}

	//发布消息
}
