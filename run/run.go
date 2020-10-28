/*
@Time : 2019-01-25 15:55
@Author : seefan
@File : run
@Software: microgo
*/
package run

import (
	"os"
	"path/filepath"
	"syscall"

	"github.com/golangteam/function/errors"
)

//Run the new Runnable
func Run(run func() Runnable, outFile ...string) {
	defer errors.PrintErr()
	cmd := "debug"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	path, err := filepath.Abs(os.Args[0])
	if err != nil {
		path = os.Args[0]
	}
	path = filepath.Dir(path)
	pidFile := filepath.Join(path, "pid.save")
	// println("cmd is ", os.Getpid(), " args is ", strings.Join(os.Args, ","))
	switch cmd {
	case "start":
		var f *os.File
		if len(outFile) > 0 {
			logFile := filepath.Join(path, outFile[0])
			tmp := filepath.Dir(logFile)
			if _, err := os.Stat(tmp); os.IsNotExist(err) {
				if err := os.MkdirAll(tmp, 0764); err != nil {
					panic(err.Error())
				}
			}
			if f, err = os.Create(logFile); err != nil {
				panic(err.Error())
			}
			defer func() {
				if err = f.Close(); err != nil {
					println(err.Error())
				}
			}()
		}
		nohup(func(sig chan<- os.Signal) {
			r := run()
			if r == nil {
				return
			}
			if err := r.Start(); err != nil {
				sig <- syscall.SIGABRT
				return
			}
			if err := r.Stop(); err != nil {
				println(err.Error())
			}

		}, f, pidFile, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	case "stop":
		if err := stopByPidFile(pidFile); err != nil {
			println(err.Error())
		}
	default:
		r := run()
		if r == nil {
			return
		}
		if err := r.Start(); err != nil {
			println(err.Error())
		}
	}
}
