package run

import (
	"io/ioutil"
	"os"
	"strconv"
)

func savePid(file string, pid int64) error {
	return ioutil.WriteFile(file, strconv.AppendInt(nil, pid, 10), 0764)
}
func getPid(file string) (int, error) {
	_, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(string(bs))
	if err != nil {
		return 0, err
	}
	return i, nil
}
