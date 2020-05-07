/*
@Time : 2019-01-24 19:28
@Author : seefan
@File : md5
@Software: le
*/
package id

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

//生成md5密码
func Md5(ss ...string) string {
	m := md5.New()
	for _, s := range ss {
		io.WriteString(m, s)
		io.WriteString(m, ":")
	}
	bs := m.Sum(nil)
	return hex.EncodeToString(bs)
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}
