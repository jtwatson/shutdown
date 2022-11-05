package shutdown_test

import (
	"context"
	"fmt"
	"syscall"
	"time"

	"github.com/jtwatson/shutdown"
)

func ExampleCaptureInterrupts() {
	// Capture interrupt so we can handle them gracefully.
	ctx, cancel := shutdown.CaptureInterrupts(context.Background())
	defer cancel()

	fmt.Println("Waiting for interrupt")

	go func() {
		time.Sleep(time.Millisecond * 100)
		_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	<-ctx.Done()

	fmt.Println("Received interrupt")
	// Output:
	// Waiting for interrupt
	// Received interrupt
}
