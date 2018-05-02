package file

import (
	"fmt"
	"os"
)

const (
	Read  = 4
	Write = 2
	Exec  = 1
)

type FileWriter struct {
	File     *os.File
	FilePath string
	IsAppend bool
}

func (f *FileWriter) WriteString(s string) error {
	if f == nil {
		return fmt.Errorf("nil pointer err")
	}
	_, err := f.File.WriteString(s)
	if err != nil {
		return err
	}
	f.File.Sync()
	return err
}

func (f *FileWriter) WriteByte(b []byte) error {
	if f == nil {
		return fmt.Errorf("nil pointer err")
	}
	_, err := f.File.Write(b)
	if err != nil {
		return nil
	}
	f.File.Sync()
	return err
}

func (f *FileWriter) Close() {
	f.File.Close()
}

//display information of the open file
func (f *FileWriter) Info() error {
	if f == nil {
		return fmt.Errorf("nil pointer err")
	}
	fileInfo, err := os.Stat(f.FilePath)
	if err != nil {
		return err
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())

	return nil
}

//rename the file and still to open it
func (f *FileWriter) Rename(newName string) error {
	if f == nil {
		return fmt.Errorf("nil pointer err")
	}

	index := -1 // '\\' is for windows
	for i := len(f.FilePath) - 1; i >= 0; i-- {
		if f.FilePath[i] == '/' || f.FilePath[i] == '\\' {
			index = i
			break
		}
	}

	newPath := f.FilePath[:index+1] + newName
	err := os.Rename(f.FilePath, newPath)
	if err != nil {
		return err
	}
	f.FilePath = newPath
	return nil
}

const (
	FileBegin = 0
	FileNow   = 1
	FileEnd   = 2
)

//seek only effect read and don`t not effect write
//but after write the seek will change
func (f *FileWriter) Seek(offset int64, whence int) error {
	if whence != FileBegin && whence != FileEnd && whence != FileNow {
		return fmt.Errorf("whence must be one of them: [FileBegin = %d, FlieNow = %d, FileEnd = %d]", FileBegin, FileNow, FileEnd)
	}

	_, err := f.File.Seek(offset, whence)
	if err != nil {
		return err
	}
	// fmt.Println("newPosition: ", newPosition)
	return nil
}

func (f *FileWriter) ReadString(length int) (string, error) {
	// read at most len(b) bytes
	// '\n' is also a byte
	b := make([]byte, length)
	_, err := f.File.Read(b)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
