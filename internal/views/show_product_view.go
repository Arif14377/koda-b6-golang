package views

import (
	"fmt"

	"github.com/Arif14377/koda-b6-golang/internal/services"
)

func MenuProduct() {
	fmt.Println("1. Show Products")
	fmt.Println("2. Search Product")
	fmt.Println("3. Filter Product")
	fmt.Println("0. Back")
}

func Product(url string) {
	products, err := services.FetchProduct(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("-- Data Produk --")
	for i := range products {

		fmt.Printf("Id: %d\n", products[i].Id)
		fmt.Printf("Name: %s\n", products[i].Name)
		fmt.Printf("Price: %d\n\n", products[i].Price)
	}
}
