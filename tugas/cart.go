package main

import (
	"context"
	"fmt"
)

func Cart(childCtx context.Context) {
	fmt.Println(childCtx.Value("cart"))
}
