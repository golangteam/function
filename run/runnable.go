/*
@Time : 2020-05-07 10:55
@Author : seefan
@File : run
@Software: microgo
*/
package run

//Runnable
type Runnable interface {
	Start() error
	Stop() error
}
