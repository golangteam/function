package errors

import (
	"path/filepath"
	"os"
	"fmt"
	"runtime"
)
// print error
//
// example: defer errors.PrintErr()
func PrintErr() {
	if err := recover(); err != nil {
		path, fe := filepath.Abs(os.Args[0])
		if fe != nil {
			path = os.Args[0]
		}
		path = filepath.Dir(path)
		path += string(os.PathSeparator) + "fault.txt"
		str := fmt.Sprintf("%v\n", err)
		for i := 1; i < 10; i++ {
			funcName, file, line, ok := runtime.Caller(i)
			if ok {
				str += fmt.Sprintf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
			}
		}
		logFile, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
		if err != nil {
			println(err.Error())
			println(str)
			return
		}
		defer logFile.Close()
		println(str)
		logFile.WriteString(str)
	}
}
