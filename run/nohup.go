/*
@Time : 2018/5/10 10:52
@Author : seefan
@File : nohup
@Software: function
*/
package run

import (
	"os"
	"os/exec"
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
		args := append([]string{}, os.Args[1:]...)
		dm := exec.Command(os.Args[0], args...)
		
		dm.Start()
		p, err := os.StartProcess(os.Args[0], args, &os.ProcAttr{Files: []*os.File{os.Stdin, out, out}})
		if err != nil {
			println(err.Error())
		} else {
			if err := savePid(pidFile, int64(p.Pid)); err != nil {
				println(err.Error())
			}
		}
		
		return
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, mon...)
	go start(sig)
	<-sig
}
