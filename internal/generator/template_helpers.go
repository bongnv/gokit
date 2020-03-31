package generator

import (
	"fmt"
	"io"
	"log"
	"strings"
	"text/template"

	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/templates"
)

func emptyValue(typeName string) string {
	if isPointer(typeName) || isArray(typeName) {
		return "nil"
	}

	if isNumber(typeName) {
		return "0"
	}

	if typeName == "string" {
		return `""`
	}

	return typeName + "{}"
}

func initValue(typeName string) string {
	if isPointer(typeName) {
		return "&" + initValue(strings.TrimPrefix(typeName, "*"))
	}

	return typeName + "{}"
}

func isPointer(typeName string) bool {
	return strings.HasPrefix(typeName, "*")
}

func isArray(typeName string) bool {
	return strings.HasPrefix(typeName, "[]")
}

func isNumber(typeName string) bool {
	return typeName == "int" ||
		typeName == "int32" ||
		typeName == "int64" ||
		typeName == "uint32" ||
		typeName == "uint64" ||
		typeName == "float64" ||
		typeName == "float32"
}

func toLower(input string) string {
	return strings.ToLower(input)
}

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func concat(inputs ...string) string {
	return strings.Join(inputs, "")
}

func toPlural(input string) string {
	return input + "s"
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"emptyValue": emptyValue,
		"isPointer":  isPointer,
		"isNumber":   isNumber,
		"replace":    replace,
		"concat":     concat,
		"toLower":    toLower,
		"initValue":  initValue,
		"toPlural":   toPlural,
	}
}

func renderFromTemplate(wr io.Writer, templateName string, data *parser.Data) error {
	log.Printf("Rendering from template %s..\n", templateName)

	templatePath := fmt.Sprintf("tmpl/%s.tmpl", templateName)
	tmplContent, err := templates.Asset(templatePath)
	if err != nil {
		return err
	}

	codeTmpl, err := template.New("default").
		Funcs(getFuncMap()).
		Parse(string(tmplContent))

	if err != nil {
		return err
	}

	if err := codeTmpl.Execute(wr, data); err != nil {
		return err
	}

	return nil
}
