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
	println("args is ", strings.Join(os.Args, " "))
	for i := 0; i < 5; i++ {
		println(i)
		time.Sleep(time.Second)
	}
	return nil
}
func (r Run) Stop() error {
	return nil
}
