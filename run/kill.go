/*
@Time : 2018/5/10 11:19 
@Author : seefan
@File : kill
@Software: function
*/
package run

import "os"

func Kill(pid int) error {
	p, err := os.FindProcess(pid)
	if err == nil {
		err = p.Kill()
		if err == nil {
			p.Wait()
		}
	}
	return err
}
