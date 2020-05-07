/*
@Time : 2018/5/10 11:19
@Author : seefan
@File : kill
@Software: function
*/
package run

import (
	"os"
	"syscall"
)

//Kill
func Kill(pid int) error {
	p, err := os.FindProcess(pid)
	if err == nil {
		err = p.Signal(syscall.SIGINT)
		if err == nil {
			for {
				if ps, e := p.Wait(); e != nil || ps != nil && ps.Exited() {
					break
				}
			}
		}
	}
	return err
}
