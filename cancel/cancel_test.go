package cancel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Cancel() error {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	go func() {

		select {
		case <-ctx.Done():
			fmt.Println("context canceled")
		case <-time.After(time.Second * 3):
			fmt.Println("job finished")
		}

	}()

	cancel()
	return ctx.Err()

}

func TestCancel(t *testing.T) {
	err := Cancel()
	if err != context.Canceled {
		t.Errorf("expected context to be canceled.")
	}
}
