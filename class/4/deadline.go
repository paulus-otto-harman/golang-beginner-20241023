package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx := context.Background()
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(parentCtx, deadline)
	defer cancel()
	result := process(ctx)
	fmt.Println("Hasil operasi:", result)
}

func process(ctx context.Context) string {
	select {
	case <-time.After(6 * time.Second):
		return "Proses selesai tepat waktu"
	case <-ctx.Done():
		// Context dibatalkan (termasuk karena deadline)
		return "Proses dibatalkan karena melewati deadline"
	}
}
