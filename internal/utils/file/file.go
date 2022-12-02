package file

import (
	"errors"
	"fmt"
	"os"

	"github.com/author_name/project_name/configs"
)

func createFile(filepath string) error {
	var _, err = os.Stat(filepath)
	if os.IsNotExist(err) {
		var file, err = os.Create(filepath)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	return nil
}

func WriteFile(filepath string, content string) error {
	fp, err := os.Create(filepath)
	if err != nil {
		return err
	}
	// write some text line-by-line to fp
	_, err = fp.WriteString(content)
	if err != nil {
		return err
	}

	// save changes
	err = fp.Sync()
	if err != nil {
		return err
	}
	return nil
}

func Mkdir(dir string) error {
	if stat, err := os.Stat(dir); err != nil {
		if err2 := os.MkdirAll(dir, 0755); err2 != nil {
			return err2
		}
	} else if !stat.IsDir() {
		return errors.New(fmt.Sprintf("%s %s is not a directory", configs.Symbol.Error, dir))
	}
	return nil
}
