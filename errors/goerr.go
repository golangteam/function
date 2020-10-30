//一个简单的生成错误的帮助类
package errors

import (
	"fmt"
)

var (
	FormatString = "%v\nthe trace error is\n%s"
)

// return a error with original error
//
// @param err     error          original error
// @param format  string         string format
// @param p       ...interface{} additional parameter
// @return error
func NewError(err error, format string, p ...interface{}) error {
	return fmt.Errorf(FormatString, fmt.Sprintf(format, p...), err)
}

// return a error
// @param format  string         string format
// @param p       ...interface{} additional parameter
// @return  error
func New(format string, p ...interface{}) error {
	return fmt.Errorf(format, p...)
}

//Returns an error in format with both the error code and the original error message
//
// @param code int the code of error
// @param format  string         string format
// @param p       ...interface{} additional parameter
// @return error
func NewCodeError(code int, format string, p ...interface{}) error {
	var txt string
	if len(p) > 0 {
		txt = fmt.Sprintf(FormatString, fmt.Sprintf(format, p...))
	} else {
		txt = format
	}
	newerr := Error{code, txt}
	return error(newerr)
}

//new error type
type Error struct {
	Code int
	Text string
}

//Implement error interface
//
// @return string
func (e Error) Error() string {
	return fmt.Sprintf("Code is %d, Error Text is %s", e.Code, e.Text)
}
