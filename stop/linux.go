package stop

import (
	"os"

	"github.com/golangteam/function/file"
	"github.com/golangteam/function/run"
)

func StopByPidFile(pidFile string) error {
	if pid, err := file.GetPid(pidFile); err == nil {
		if err = run.Kill(pid); err != nil {
			return err
		} else {
			return os.Remove(pidFile)
		}
	} else {
		return err
	}
}
