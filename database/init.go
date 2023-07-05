package database

import "github.com/xiaolibuzai-ovo/L-gateway/database/mysql"

func Init() {
	mysql.InitMysql()
}
