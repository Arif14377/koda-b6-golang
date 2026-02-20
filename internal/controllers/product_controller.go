package controllers

import (
	"fmt"
	"os"

	"github.com/Arif14377/koda-b6-golang/internal/models"
	"github.com/Arif14377/koda-b6-golang/internal/services"
	"github.com/Arif14377/koda-b6-golang/internal/views"
)

var url string = "https://raw.githubusercontent.com/Arif14377/koda-b6-weekly1/refs/heads/main/data.json"
var cart models.Cart

func StartAppController() {
	views.Landing()
	menu := views.InputAngka("Input menu: ")

	switch menu {
	case 1:
		LoginController()
	case 2:
		MenuProductController(url)
	}
}

func MenuProductController(url string) {
	views.MenuProduct()
	menu := views.InputAngka("Input menu: ")

	switch menu {
	case 1:
		ShowProductController(url, cart)
	case 2:
		SearchProductController(url, cart)
	case 3:
		FilterProductController()
	case 0:
		StartAppController()
	}
}

func ShowProductController(url string, cart models.Cart) {
	views.Product(url)
	for true {
		fmt.Println("# What do you want:")
		fmt.Println("1. Masukkan ke keranjang")
		fmt.Println("2. Checkout")
		pilihan := views.InputAngka("Masukkan angka pilihan: ")
		// var cart []models.CartItem

		// 1. keranjang; 2. checkout.
		switch pilihan {
		case 1:
			id := views.InputAngka("Masukkan id product: ")
			qty := views.InputAngka("Masukkan jumlah produk: ")

			cart.Cart = append(cart.Cart, services.AddCart(id, qty, url))
			fmt.Printf("Cart Anda: %v\n", cart.Cart)

			totalPrice := services.TotalPrice(cart.Cart)
			fmt.Printf("Total Belanja: Rp. %d\n", totalPrice)

			continue
		}
	}
}

func SearchProductController(url string, cart models.Cart) {
	searchKeyword := views.InputString("Search keyword: ")

	products, _ := services.FetchProduct(url)

	result, err := services.SearchProduct(searchKeyword, products)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	os.Exit(0)
}

func FilterProductController() {

}
