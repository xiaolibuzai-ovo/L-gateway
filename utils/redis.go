package utils

import (
	"github.com/gomodule/redigo/redis"
	db_redis "github.com/xiaolibuzai-ovo/L-gateway/database/redis"
)

func RedisBatchOperate(functions ...func(c redis.Conn)) error {
	redisConn := db_redis.GetRedisConn()
	defer redisConn.Close()
	for _, f := range functions {
		f(redisConn)
	}
	redisConn.Flush()
	return nil
}

func RedisConfDo(commandName string, args ...interface{}) (interface{}, error) {
	redisConn := db_redis.GetRedisConn()
	defer redisConn.Close()
	return redisConn.Do(commandName, args...)
}
