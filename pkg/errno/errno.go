package errno

import "fmt"

// 错误码设计
type Errno struct {
	Code     int
	Msg      string
	Internal error // 内部错误
}

// 实现了 error 接口
func (e *Errno) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s, 内部错误信息: %s", e.Code, e.Msg, e.Internal.Error())
}

// 定义新的错误码
func NewErrno(code int, msg string, err error) *Errno {
	return &Errno{
		Code:     code,
		Msg:      msg,
		Internal: err,
	}
}

func New(errno *Errno, err error) *Errno {
	return &Errno{
		Code:     errno.Code,
		Msg:      errno.Msg,
		Internal: err,
	}
}

// 解码错误, 返回 code 和 msg
func DecodeErr(err error) (int, string) {
	if err == nil {
		return Ok.Code, Ok.Msg
	}
	switch t := err.(type) {
	case *Errno:
		return t.Code, t.Msg
	default:
		return UNKNOWN.Code, err.Error()
	}
}
