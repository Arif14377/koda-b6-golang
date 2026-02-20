package models

type Cart struct {
	Cart  []CartItem
	Total int
}

type CartItem struct {
	Id    int
	Name  string
	Price int
	Qty   int
}

func (c *CartItem) SumPriceItem() int {
	return c.Price * c.Qty
}
