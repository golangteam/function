//一个简单的生成错误的帮助类
package errors

import (
	"fmt"
)

var (
	FormatString = "%v\nthe trace error is\n%s"
)

//按格式返回一个错误，同时携带原始的错误信息
//
//  返回 error
func NewError(err error, format string, p ...interface{}) error {
	return fmt.Errorf(FormatString, fmt.Sprintf(format, p...), err)
}

//返回一个错误
//
//  返回 error
func New(format string, p ...interface{}) error {
	return fmt.Errorf(format, p...)
}

//按格式返回一个错误，同时携带错误代码和原始的错误信息
//
//  返回 error
func NewCodeError(code int, format string, p ...interface{}) error {
	txt := fmt.Sprintf(FormatString, fmt.Sprintf(format, p...))
	newerr := Error{code, txt}
	return error(newerr)
}

//新的error类型
type Error struct {
	Code int
	Text string
}

//实现error接口
func (this Error) Error() string {
	return fmt.Sprintf("Code is %d, Error Text is %s", this.Code, this.Text)
}
