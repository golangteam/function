package run

import (
	"os"
)

func stopByPidFile(pidFile string) error {
	if pid, err := getPid(pidFile); err == nil {
		if err = Kill(pid); err != nil {
			return err
		}
		return os.Remove(pidFile)
	} else {
		return err
	}
}
