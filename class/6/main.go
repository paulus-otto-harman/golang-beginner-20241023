package main

import (
	"context"
	"fmt"
	"time"
)

func satu(ctx context.Context, key interface{}, val interface{}) context.Context {
	return context.WithValue(ctx, key, val)
}

func dua(ctx context.Context, key interface{}) {
	nilai := ctx.Value(key)
	for {

		fmt.Println(nilai)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	parentCtx := context.Background()
	type key string

	childCtx := satu(parentCtx, key("satu"), "satu")

	fmt.Println(childCtx.Value(key("satu")))

	go dua(childCtx, key("satu"))

	time.Sleep(time.Second * 7)

	childCtx = satu(parentCtx, key("satu"), "dua")
	time.Sleep(time.Second * 7)
}
