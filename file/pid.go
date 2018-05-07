package file

import (
	"io/ioutil"
	"os"
	"strconv"
)

func SavePid(file string) error {
	return ioutil.WriteFile(file, strconv.AppendInt(nil, int64(os.Getpid()), 10), 0764)
}
func GetPid(file string) (string, error) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
