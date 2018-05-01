package file

import (
	"fmt"
	"testing"
)

var (
	s = "test_string\n"
	b = []byte("test_byte\n") // not a const
)

const (
	filePath = "./test.txt"
)

func checkerr(t *testing.T, err error) {
	if err != nil {
		fmt.Println(err.Error())
		t.Fatal(err)
	}
}

func TestWriteAppend(t *testing.T) {
	f, err := NewFileWriter(filePath, true)
	checkerr(t, err)
	defer f.Close()

	err = f.WriteString(s)
	checkerr(t, err)

	err = f.WriteByte(b)
	checkerr(t, err)
}

func TestWriteNotAppend(t *testing.T) {
	f, err := NewFileWriter(filePath, false)
	checkerr(t, err)
	defer f.Close()

	err = f.WriteString(s)
	checkerr(t, err)

	err = f.WriteByte(b)
	checkerr(t, err)
}

func TestFileInfo(t *testing.T) {
	f, err := NewFileWriter(filePath, true)
	checkerr(t, err)
	defer f.Close()

	err = f.Info()
	checkerr(t, err)
}

// // func TestReaname(t *testing.T) {
// // 	f, err := NewFileWriter(filePath, false)
// // 	checkerr(t, err)
// // 	defer f.Close()

// // 	err = f.Rename("vimi.txt")
// // 	checkerr(t, err)

// // 	err = f.WriteString("haha\n")
// // 	checkerr(t, err)

// // 	err = f.WriteString("haha")
// // 	checkerr(t, err)
// // }

func TestFileExist(t *testing.T) {
	if FileExist("no") == true {
		checkerr(t, fmt.Errorf("bad expection"))
	}

	if FileExist(filePath) == false {
		checkerr(t, fmt.Errorf("bad expection"))
	}
}

func TestSeek(t *testing.T) {
	f, err := NewFileWriter(filePath, true)
	checkerr(t, err)
	defer f.Close()

	err = f.Seek(1, FileBegin)
	checkerr(t, err)

}

func TestReadString(t *testing.T) {
	f, err := NewFileWriter(filePath, true)
	checkerr(t, err)
	defer f.Close()

	res, err := f.ReadString(40)
	checkerr(t, err)
	fmt.Println(res)
}
