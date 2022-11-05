# shutdown

[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/jtwatson/shutdown)

**shutdown** is a simple package that helps facilitate a graceful shutdown of your application. It can create a context that will be canceled when your program receives SIGTERM or SIGINT from the os.

## Installation

    go get github.com/jtwatson/shutdown@latest

## Simple Example

```go
package main

import (
	"context"
	"fmt"
	"syscall"
	"time"

	"github.com/jtwatson/shutdown"
)

func main() {
	// Capture interrupt so we can handle them gracefully.
	ctx, cancel := shutdown.CaptureInterrupts(context.Background())
	defer cancel()

	fmt.Println("Waiting for interrupt")

	go func() {
		time.Sleep(time.Second)
		_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	<-ctx.Done()

	fmt.Println("Received interrupt")
}
```
