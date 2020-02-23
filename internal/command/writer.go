package command

import (
	"os"
	"path/filepath"
)

type writer interface {
	Write(path string, data []byte) error
}

type fileWriter struct{}

func (fw *fileWriter) Write(path string, data []byte) error {
	if err := fw.ensurePath(filepath.Dir(path)); err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func() {
		// TODO: write into log
		_ = f.Close()
	}()

	_, err = f.Write(data)
	return err
}

func (fw *fileWriter) ensurePath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}

	return nil
}
