/*
	config 定义配置项
*/
package configs

import "time"

// 时区
var TimeZone *time.Location

// mysql 连接
var MysqlDsn string = "root:1234@tcp(127.0.0.1:3306)/xingkong?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	TimeZone, _ = time.LoadLocation("Asia/Shanghai")
}
