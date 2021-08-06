/*
	errno 定义了错误码

	错误码设计为 5 位长度
	第一位是错误级别, 1 为系统错误, 2 为用户错误
	第二三位是模块编号
	第四五位是具体错误编号

	包含两个特例.
	0 表示请求成功, 没有发生任何错误.
	1 表示错误没有编号, 属于未知错误, 或服务器内部错误.
*/
package errno

// 定义必备错误码
var (
	Ok      = &Errno{Code: 0, Msg: "ok"}
	UNKNOWN = &Errno{Code: 1, Msg: "unknown"}
)

// 定义错误码, 以及默认的错误信息

// 1 系统错误

// 2 用户错误

// 201 请求相关
var ErrBind = &Errno{Code: 20101, Msg: "请求参数错误"}

// 202 数据库
var ErrNotFound = &Errno{Code: 20201, Msg: "数据不存在"}
