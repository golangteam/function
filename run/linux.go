package run

import (
	"os"
)

func stopByPidFile(pidFile string) error {
	pid, err := getPid(pidFile)
	if err == nil {
		if err = Kill(pid); err != nil {
			return err
		}
		return os.Remove(pidFile)

	}
	return err
}
