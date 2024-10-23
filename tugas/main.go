package main

import (
	"context"
	"time"
)

func main() {
	const SessionTimeout = 3
	parentCtx := context.Background()

	ctx, cancel := context.WithTimeout(parentCtx, SessionTimeout*time.Second)
	defer cancel()

	go navigate(ctx)
}

func navigate(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
