package main

import (
	"fmt"
	"goRedisDemo/handler"
)

func main() {
	// fmt.Println("test Redis")
	// handler.UploadHandler()

	fmt.Println("test mySQL ")
	handler.UploadHandlerDB()
}
