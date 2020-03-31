package generator

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"

	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/iohelper"
	"golang.org/x/tools/imports"
)

// Appender renders a template and appends to an existing file.
type Appender struct {
	FilePath     string
	TemplateName string
	Data         *parser.Data
	Writer       iohelper.Writer
	Reader       iohelper.Reader
}

// Do appends content to the given file given a template.
func (a *Appender) Do() error {
	existingData, err := a.Reader.Read(a.FilePath)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(existingData)
	buf.WriteString("\n")

	err = renderFromTemplate(buf, a.TemplateName, a.Data)
	if err != nil {
		return err
	}

	log.Println("Formatting ", filepath.Base(a.FilePath))
	formatedBuf, err := imports.Process(a.FilePath, buf.Bytes(), nil)
	if err != nil {
		fmt.Println("Failed to format. Content:", buf.String())
		return err
	}

	return a.Writer.Write(a.FilePath, formatedBuf)
}
