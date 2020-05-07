/*
@Time : 2019-01-20 20:46
@Author : seefan
@File : number
@Software: le
*/
package id

import (
	"fmt"
	"math/rand"
	"strconv"
)

//判断参数是否可转化为数字
func IsNumber(src interface{}) bool {
	switch v := src.(type) {
	case int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
		return true
	default:
		word := fmt.Sprint(v)
		if word[0] >= 48 && word[0] <= 57 {
			return true
		}
	}
	return false
}

//返回一个随机的数字
func RndString() string {
	return strconv.FormatInt(rand.Int63(), 16)
}
