package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func ensureDir(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.FileMode(0777))
		if merr != nil {
			panic(merr)
		}
	}
}

func CreateFilesIfNotExists(filenames []string) error {
	for _, filename := range filenames {
		if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
			ensureDir(filename)
			file, err := os.Create(filename)
			if err != nil {
				return err
			}
			defer file.Close()
		}
	}
	return nil
}
