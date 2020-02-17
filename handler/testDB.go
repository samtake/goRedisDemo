package handler

import (
	mydb "goRedisDemo/db"
	"strconv"
)

func UploadHandlerDB() {
	fsize, _ := strconv.Atoi("filesize")
	mydb.OnFileUploadFinished("testHash", "testName", int64(fsize), "")
}
