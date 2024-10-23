package main

import (
	"context"
	"fmt"
)

func satu(ctx context.Context, k interface{}) {
	fmt.Println(ctx.Value(k))
}

func dua(ctx context.Context, k interface{}) {
	ctxValue := ctx.Value(k)
	angka, ok := ctxValue.(int)
	if ok {
		fmt.Println(angka + 5)
	}

}

func main() {
	parentCtx := context.Background()
	type key string

	ctx1 := context.WithValue(parentCtx, key("nilai"), 10)

	satu(ctx1, key("nilai"))
	dua(ctx1, key("nilai"))
}
