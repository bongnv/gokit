package command

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"text/template"

	"github.com/bongnv/gokit/internal/templates"
	"golang.org/x/tools/imports"
)

type fileGenerator struct {
	filePath     string
	templateName string
	service      *Service
	writer       writer
}

// Do generates a file given a template. Then, it uses writer to render output.
func (g *fileGenerator) Do() error {
	buf, err := g.renderFromTemplate()
	if err != nil {
		return err
	}

	log.Println("Formatting ", filepath.Base(g.filePath))
	buf, err = imports.Process(g.filePath, buf, nil)
	return g.writer.Write(g.filePath, buf)
}

func (g *fileGenerator) renderFromTemplate() ([]byte, error) {
	log.Printf("Rendering from template %s..\n", endpointsTemplateName)

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
