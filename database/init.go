package database

import (
	"github.com/xiaolibuzai-ovo/L-gateway/database/mysql"
	"github.com/xiaolibuzai-ovo/L-gateway/database/redis"
)

func Init() {
	db_mysql.InitMysql()
	db_redis.InitRedis()
}
