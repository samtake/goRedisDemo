package handler

import (
	"encoding/json"
	"filestore-server/mq"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	cfg "goRedisDemo/config"
)

// UploadHandler ： 
func UploadHandler() {
	// 写入异步转移任务队列
	data := mq.TransferData{
		FileHash:      fileMeta.FileSha1,
		CurLocation:   fileMeta.Location,
		DestLocation:  ossPath,
		DestStoreType: cmn.StoreOSS,
	}
	pubData, _ := json.Marshal(data)
	pubSuc := mq.Publish(
		cfg.TransExchangeName,
		cfg.TransOSSRoutingKey,
		pubData,
	)
	if !pubSuc {
		// TODO: 当前发送转移信息失败，稍后重试
	}
}
