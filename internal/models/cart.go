package models

type Cart struct {
	Cart  []CartItem
	Total int
}

type CartItem struct {
	Product Product
	Qty     int
}
