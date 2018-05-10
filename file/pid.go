package file

import (
	"io/ioutil"
	"os"
	"strconv"
)

func SavePid(file string) error {
	return ioutil.WriteFile(file, strconv.AppendInt(nil, int64(os.Getpid()), 10), 0764)
}
func GetPid(file string) (int, error) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, err
	}
	if i, err := strconv.Atoi(string(bs)); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}
