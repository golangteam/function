/*
@Time : 2018/5/10 10:52 
@Author : seefan
@File : nohup
@Software: function
*/
package run

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/golangteam/function/errors"
)

//nohup function
func Nohup(start func() error, exit func(error), out *os.File, mon ...os.Signal) {
	if os.Getppid() != 1 {
		args := append([]string{os.Args[0]}, os.Args[1:]...)
		if out == nil {
			out = os.NewFile(uintptr(syscall.Stderr), "/dev/null")
			defer out.Close()
		}
		if _, err := os.StartProcess(os.Args[0], args, &os.ProcAttr{Files: []*os.File{os.Stdin, out, out}}); err != nil {
			println(err.Error())
		}
		return
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, mon...)
	go func() {
		if err := start(); err != nil {
			if exit != nil {
				exit(err)
			}
			sig <- syscall.SIGABRT
		} else {
			sig <- syscall.SIGQUIT
		}
	}()
	s := <-sig
	if s != syscall.SIGABRT && exit != nil {
		exit(errors.New("signal:", s))
	}
}
