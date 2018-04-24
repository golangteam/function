package command

import (
	"io/ioutil"
	"os/exec"
)

/**
 * [Command exec some terminal command and read result]
 *
 * @param {[type]} command string) (res string, err error)
 */
func Command(command string) (res string, err error) {
	cmd := exec.Command(command)

	stdout, err := cmd.StdoutPipe()
	defer stdout.Close()

	if err != nil {
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	return string(opBytes), nil
}
