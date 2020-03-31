package iohelper

import (
	"os"
	"path/filepath"
)

// Writer is an interface to wrap Write method.
//go:generate mockery -name=Writer -inpkg -case=underscore
type Writer interface {
	Write(path string, data []byte) error
}

// FileWriter is an implementation of Writer to file to file.
type FileWriter struct{}

func (fw FileWriter) Write(path string, data []byte) error {
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

func (fw FileWriter) ensurePath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}

	return nil
}
