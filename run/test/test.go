package main

import (
	"os"
	"strings"
	"time"

	"github.com/golangteam/function/run"
)

func main() {
	run.Run(func() run.Runnable {
		return new(Run)
	})
}

type Run struct{}

func (r Run) Start() error {
	println("run start")
	println("args is ", strings.Join(os.Args, " "))
	for i := 0; i < 2; i++ {
		println(i)
		time.Sleep(time.Second)
	}
	time.Sleep(time.Minute)
	return nil
}
func (r Run) Stop() error {
	println("run stop")
	return nil
}
