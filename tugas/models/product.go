package models

type Product struct {
	name string
}

func InitProduct(name string) Product {
	return Product{name}
}
