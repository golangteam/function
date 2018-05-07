package stop

import (
	"os/exec"
	"time"
)

func LinuxPid(pid string, sig ...string) (err error) {
	if len(sig) == 0 {
		sig = append(sig, "USR2")
	}

	checkCmd := exec.Command("kill", "-s", "0", pid)
	killCmd := exec.Command("kill", "-s", sig[0], pid)
	now := time.Now()
	err = killCmd.Run()
	if err == nil {
		for {
			if err := checkCmd.Run(); err != nil {
				break
			}
			time.Sleep(time.Millisecond * 100)
			if time.Since(now).Seconds() > 30 {
				break
			}
		}
	}
	return
}
