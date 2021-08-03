/*
	config 定义配置项
*/
package configs

import "time"

// 时区
var TimeZone *time.Location

func init() {
	TimeZone, _ = time.LoadLocation("Asia/Shanghai")
}
