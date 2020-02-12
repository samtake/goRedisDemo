package handler

import (
	rPool "goRedisDemo/cache/redis"
)

func UploadHandler() {
	// 获得redis的一个连接
	rConn := rPool.RedisPool().Get()
	defer rConn.Close()

	// 将初信息写入到redis缓存
	rConn.Do("sadd", "set", "uploadHandler")
}
