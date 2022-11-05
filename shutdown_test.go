package shutdown

import (
	"context"
	"syscall"
	"testing"
	"time"
)

func TestCaptureInterruptsWithCancel(t *testing.T) {
	tests := []struct {
		name       string
		cancelFunc func(parent, child context.CancelFunc)
	}{
		{name: "cancel parent context", cancelFunc: func(parent, child context.CancelFunc) { parent() }},
		{name: "cancel child context ", cancelFunc: func(parent, child context.CancelFunc) { child() }},
		{name: "send SIGTERM", cancelFunc: func(parent, child context.CancelFunc) { _ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM) }},
		{name: "send SIGINT ", cancelFunc: func(parent, child context.CancelFunc) { _ = syscall.Kill(syscall.Getpid(), syscall.SIGINT) }},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			parent, parentCancel := context.WithCancel(context.Background())
			gotCtx, gotCancel := CaptureInterrupts(parent)

			tt.cancelFunc(parentCancel, gotCancel)

			select {
			case <-gotCtx.Done():
				// Allow time for goroutine to exit so we can observe code coverage within go-routine.
				// It isn't possible (or necessary) with current design to assert that the go routine exists.
				time.Sleep(time.Millisecond * 100)
			case <-time.After(time.Second * 5):
				t.Error("CaptureInterruptsWithCancel() context should have been canceled, but was not")
			}
		})
	}
}
