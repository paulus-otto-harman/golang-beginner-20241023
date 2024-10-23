package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx := context.Background()

	ctx, cancel := context.WithTimeout(parentCtx, 6*time.Second)
	defer cancel()

	angka := 10

	go process(ctx, angka)
	time.Sleep(time.Second * 10)
}

func process(ctx context.Context, nilai int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			nilai--
			fmt.Println(nilai)
		}
		time.Sleep(time.Second * 2)
	}

}
