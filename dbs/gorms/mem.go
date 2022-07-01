package gorms

import (
	"github.com/xuender/oils/base"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewMemDB() *gorm.DB {
	memDB := base.Must1(gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   NewGormLogger(),
	}))

	memDB.Debug()

	return memDB
}
