# Fiber redoc

[![Build Status](https://travis-ci.org/ivpusic/grpool.svg?branch=master)](https://github.com/infinitasx/easi-go-aws)

## ⚙ Installation

```text
go get -u github.com/eininst/fiber-middleware-redoc
```

## ⚡ Quickstart

```go
package main

import (
	redoc "github.com/eininst/fiber-middleware-redoc"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/redoc/*", redoc.New("examples/swagger.json"))

    app.Listen(":8080")
}
```
> to http://127.0.0.1:8080/redoc
> 
 <img alt="Redoc logo" src="https://fab-jar.oss-cn-zhangjiakou.aliyuncs.com/img/redoc.png"  width="640px"/>


> See [examples](/examples)

## License

*MIT*