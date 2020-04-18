# Getting Started

The quick start creates a Hello World service using [`gokitgen`](https://github.com/bongnv/gokit).

## Install `gokitgen`

Make sure [Golang](https://golang.org/doc/install) is installed. Use the following command to install `gokitgen` from source:

```bash
go get -u github.com/bongnv/gokit/cmd/gokitgen
```

## Scaffold a project

```bash
mkdir hello
cd hello
```

Use `gokitgen` to scaffold an empty project:
```bash
gokitgen scaffold -package github.com/hello
```

By default, `service.go` will be generated with two example endpoints.

```go
//go:generate gokitgen service -interface Service
type Service interface {
	Hello(ctx context.Context, p *Request) (*Response, error)
	Bye(ctx context.Context, req *ByeRequest) (*ByeResponse, error)
}
```
