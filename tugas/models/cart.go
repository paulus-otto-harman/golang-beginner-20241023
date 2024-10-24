package models

type Cart struct {
	items []Product
}

func (cart *Cart) AddItem(item Product) {
	cart.items = append(cart.items, item)
}
