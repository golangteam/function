/*
@Time : 2019-01-25 15:55
@Author : seefan
@File : run
@Software: microgo
*/
package run

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/golangteam/function/errors"
)

//ParseFlag parse flags
func ParseFlag(args []string) map[string]string {
	re := make(map[string]string)
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			as := strings.Split(arg, "=")
			if len(as) == 1 {
				re[strings.TrimSpace(string(as[0][1:]))] = ""
			} else if len(as) == 2 {
				re[strings.TrimSpace(string(as[0][1:]))] = strings.TrimSpace(as[1])
			}
		} else {
			re[arg] = ""
		}
	}
	return re
}

//PrintUseage print help infomation
func PrintUseage(name string, arg map[string]string) {
	if _, ok := arg["-help"]; ok {
		fmt.Printf("Useage: %s start -conf=prod.yaml\n", name)
		fmt.Printf("        %s stop", name)
		os.Exit(0)
	}
	if _, ok := arg["h"]; ok {
		fmt.Printf("Useage: %s start -conf=prod.yaml", name)
		fmt.Printf("        %s stop", name)
		os.Exit(0)
	}
}

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
		nohup(func(sig chan os.Signal, exit chan os.Signal) {
			r := run()
			if r == nil {
				println("Runnable stoped")
				exit <- syscall.SIGABRT
				return
			}
			go func() {
				if err := r.Start(); err != nil {
					sig <- syscall.SIGABRT
				}
			}()
			<-sig //等待结束信号
			if err := r.Stop(); err != nil {
				println(err.Error())
				exit <- syscall.SIGABRT
			} else {
				exit <- syscall.SIGQUIT
			}

		}, f, pidFile, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	case "stop":
		if err := StopByPidFile(pidFile); err != nil {
			println(err.Error())
		}
	default:
		r := run()
		if r == nil {
			return
		}
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		go func() {
			if err := r.Start(); err != nil {
				sig <- syscall.SIGABRT
			}
		}()
		<-sig //等待结束信号
		if err := r.Stop(); err != nil {
			println(err.Error())
		}
	}
}
