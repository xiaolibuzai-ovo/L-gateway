package dal

import "L-gateway/biz/dal/mysql"

func Init() {
	mysql.InitMysql()
}
