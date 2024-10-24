package main

import (
	"20241023/tugas/models"
	"context"
	"fmt"
)

func Products(childCtx context.Context) {
	fmt.Println("Products")
	fmt.Println("[1] Satu")
	fmt.Println("[2] Dua")
	fmt.Println("[3] Tiga")

	menuItem, _ := Input(Args(P("type", "number")))
	switch menuItem {
	case 1:
		product := models.InitProduct("Satu")
		products := childCtx.Value("cart")
		if products == nil {
			products = products.([]models.Product)
			//products = append(products, product)
		}
		context.WithValue(childCtx, "cart", product)
	case 2:
		product := models.InitProduct("Dua")
		context.WithValue(childCtx, "cart", product)
	case 3:
		product := models.InitProduct("Tiga")
		context.WithValue(childCtx, "cart", product)
	}
}
