package main

import (
	"context"
	"fmt"
	"time"
)

type Angka struct {
	nilai int
}

func main() {
	parentCtx := context.Background()
	deadline := time.Now().Add(4 * time.Second)
	ctx, cancel := context.WithDeadline(parentCtx, deadline)
	defer cancel()

	process(ctx)
}

func process(ctx context.Context) {
	var daftar []Angka
	for {
		select {
		case <-ctx.Done():
			fmt.Println("deadline exceeded")
			fmt.Println(daftar)
			return
		default:
			daftar = append(daftar, Angka{nilai: len(daftar) + 1})
			time.Sleep(time.Second)
			fmt.Println(daftar)
		}
	}
}
