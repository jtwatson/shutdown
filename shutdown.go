// shutdown is a simple package for creating a context that will be canceled
// if your program receives SIGTERM or SIGINT from the os
package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// CaptureInterrupts is the same as context.WithCancel, but the returned context will cancel if your program
// receives an interrupt from the os
func CaptureInterrupts(parent context.Context) (ctx context.Context, cancel context.CancelFunc) {
	c, cancel := context.WithCancel(parent)

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		defer signal.Stop(interruptChan)

		select {
		case <-c.Done():
			cancel()
		case <-interruptChan:
			cancel()
		}
	}()

	return c, cancel
}
