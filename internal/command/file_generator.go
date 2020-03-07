package command

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"text/template"

	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/templates"
	"github.com/bongnv/gokit/internal/writer"
	"golang.org/x/tools/imports"
)

type fileGenerator struct {
	filePath     string
	templateName string
	service      *parser.Service
	writer       writer.Writer
}

// Do generates a file given a template. Then, it uses writer to render output.
func (g *fileGenerator) Do() error {
	buf, err := g.renderFromTemplate()
	if err != nil {
		return err
	}

	log.Println("Formatting ", filepath.Base(g.filePath))
	formatedBuf, err := imports.Process(g.filePath, buf, nil)
	if err != nil {
		fmt.Println("Failed to format. Content:", string(buf))
		return err
	}

	return g.writer.Write(g.filePath, formatedBuf)
}

func (g *fileGenerator) renderFromTemplate() ([]byte, error) {
	log.Printf("Rendering from template %s..\n", g.templateName)

	templatePath := fmt.Sprintf("tmpl/%s.tmpl", g.templateName)
	tmplContent, err := templates.Asset(templatePath)
	if err != nil {
		return nil, err
	}

	codeTmpl, err := template.New("default").
		Funcs(getFuncMap()).
		Parse(string(tmplContent))

	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if err := codeTmpl.Execute(buf, g.service); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
