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
	"path/filepath"
	"syscall"
)

//nohup function
func nohup(start func(chan<- os.Signal), out *os.File, pidFile string, mon ...os.Signal) {
	//println("nonup args is ", strings.Join(os.Args, " "))
	if os.Getppid() != 1 {
		if out == nil {
			out = os.NewFile(uintptr(syscall.Stderr), "/dev/null")
			defer func() {
				if err := out.Close(); err != nil {
					println(err.Error())
				}
			}()
		}
		file := os.Args[0]
		if f, err := filepath.Abs(os.Args[0]); err == nil {
			file = f
		}
		_, err := os.StartProcess(file, os.Args, &os.ProcAttr{Files: []*os.File{os.Stdin, out, out}})
		if err != nil {
			println(err.Error())
		}
		return
	}
	if err := savePid(pidFile, int64(os.Getpid())); err != nil {
		println(err.Error())
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, mon...)
	go start(sig)
	<-sig
}
