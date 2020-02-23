package command

import (
	"strings"
	"text/template"
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

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func concat(inputs ...string) string {
	return strings.Join(inputs, "")
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"emptyValue": emptyValue,
		"isPointer":  isPointer,
		"isNumber":   isNumber,
		"replace":    replace,
		"concat":     concat,
	}
}
