package model

import (
	"github.com/zhenhua32/xingkong/configs"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// DB 是 GORM 数据库连接
var DB *gorm.DB

// 初始化数据库连接
func InitDB() {
	var err error
	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault()
	DB, err = gorm.Open(mysql.Open(configs.MysqlDsn), &gorm.Config{
		PrepareStmt: true,
		// Logger:      logger.Default.LogMode(logger.Silent),  // 禁用 log
		Logger: logger,
	})
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
