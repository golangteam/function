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
)

//nohup function
func nohup(start func(chan<- os.Signal), out *os.File, pidFile string, mon ...os.Signal) {
	if os.Getppid() != 1 {
		if out == nil {
			out = os.NewFile(uintptr(syscall.Stderr), "/dev/null")
			defer func() {
				if err := out.Close(); err != nil {
					println(err.Error())
				}
			}()
		}
		//println("args is ", strings.Join(os.Args, " "))
		_, err := os.StartProcess(os.Args[0], os.Args, &os.ProcAttr{Files: []*os.File{os.Stdin, out, out}})
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
