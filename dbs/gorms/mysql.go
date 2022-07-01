package gorms

import (
	"time"

	"github.com/xuender/oils/base"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB(dsn string) *gorm.DB {
	mysqlDB := base.Must1(gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   NewGormLogger(),
	}))
	sqlDB := base.Must1(mysqlDB.DB())
	two := 200
	one := 100

	sqlDB.SetMaxOpenConns(two)
	sqlDB.SetMaxIdleConns(one)
	sqlDB.SetConnMaxLifetime(time.Minute)

	return mysqlDB
}
