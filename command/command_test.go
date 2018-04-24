package command

import (
	"fmt"
	"testing"
)

const (
	PWD  = "pwd"
	LS   = "ls"
	DATE = "date"
)

func TestCommand(t *testing.T) {
	result, err := Command(PWD)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(PWD)
	fmt.Println(result)

	result, err = Command(LS)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(LS)
	fmt.Println(result)

	result, err = Command(DATE)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(DATE)
	fmt.Println(result)
}

/*
pwd
/usr/lib/go-1.9/GOPATH/src/github.com/function/command

ls
command.go
command_test.go

date
2018年 04月 24日 星期二 20:17:09 CST

PASS
ok  	github.com/function/command	0.005s
*/
