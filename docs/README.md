# Introduction

[go-kit](https://github.com/go-kit/kit) is a powerful framework to build microservices. However, it requires lots of boilerplate codes; hence it can introduce unnecessary complexibity in codebase.

This tool is created to generate those boilerplate codes, thus to avoid potential human mistake when writing them. Therefore, the tool also introduced an opinioned implementation of `go-kit`.

## How it works

The tool uses [`packages.Load`](golang.org/x/tools/go/packages) to retrieve information of an interface or struct. Then it leverages [go template](https://golang.org/pkg/text/template/) together with parsed information to generate codes from an interface or struct definition.

In order to reduce the complexity in templates, logic is prefered to be implemented in a [util](https://github.com/bongnv/gokit/tree/master/util) package.

## Features

* [Scaffolding projects](./scaffolding-projects.md)
* [Scaffolding CRUD endpoints](./scaffolding-crud-endpoints.md)
* [Generating services](./generating-services.md)
* [Generating DAO](./generating-dao.md)
