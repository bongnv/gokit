package gokit

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/bongnv/gokit/internal/templates"
	"golang.org/x/tools/imports"
)

const (
	endpointsTemplateName = "endpoints"
	endpointsFileName     = "z_endpoints.go"
)

func (h *handler) generateEndpoints() error {
	buf, err := h.renderFromTemplate(endpointsTemplateName)
	if err != nil {
		return err
	}

	endpointFile := path.Join(h.opts.Path, "internal", "endpoint", endpointsFileName)

	log.Println("Formatting ...")
	buf, err = imports.Process(endpointFile, buf, nil)
	return h.writeToFile(endpointFile, buf)
}

func (h *handler) writeToFile(path string, data []byte) error {
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

func (h *handler) renderFromTemplate(templateName string) ([]byte, error) {
	log.Printf("Rendering from template %s..\n", endpointsTemplateName)

	templatePath := fmt.Sprintf("tmpl/%s.tmpl", templateName)
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
	if err := codeTmpl.Execute(buf, h.service); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
