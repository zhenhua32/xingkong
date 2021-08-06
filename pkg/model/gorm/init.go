package model

import (
	"github.com/zhenhua32/xingkong/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 是 GORM 数据库连接
var DB *gorm.DB

// 初始化数据库连接
func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(configs.MysqlDsn), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库")
	}
}

// SetUpDB 迁移数据库
func SetUpDB() {
	var err error
	err = DB.AutoMigrate(&Book{})
	if err != nil {
		panic("无法迁移数据库 BookModel")
	}
	err = DB.AutoMigrate(&Chapter{})
	if err != nil {
		panic("无法迁移数据库 ChapterModel")
	}
}
