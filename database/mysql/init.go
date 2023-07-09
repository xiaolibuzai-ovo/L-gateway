package db_mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

// todo 可配置化

func InitMysql() {
	var dsn = "root:lmz0521@tcp(123.57.130.104:3306)/L-gateway_1?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}

func GetMysqlConn() (*gorm.DB, error) {
	return DB, err
}
