package services

import (
	"fmt"

	"github.com/Arif14377/koda-b6-golang/internal/models"
)

func AddCart(id int, qty int, url string) models.CartItem {
	var product models.CartItem

	products, err := FetchProduct(url)
	if err != nil {
		fmt.Println(err)
	}

	for i := range products {
		if products[i].Id == id {
			product = models.CartItem{
				Id:    id,
				Name:  products[i].Name,
				Price: products[i].Price * qty,
				Qty:   qty,
			}
		}
	}

	return product
}

func TotalPrice(cart []models.CartItem) int {
	totalPrice := 0
	for x := range cart {
		totalPrice = totalPrice + cart[x].Price
	}
	return totalPrice
}
