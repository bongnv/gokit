package iohelper

import "io/ioutil"

// Reader is the interface to wrap Read method.
//go:generate mockery -name=Reader -inpkg -case=underscore
type Reader interface {
	Read(path string) ([]byte, error)
}

// FileReader implements Reader interface. It allows to read a file.
type FileReader struct{}

func (fr FileReader) Read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
