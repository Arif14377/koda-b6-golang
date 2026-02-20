package controllers

import (
	"fmt"

	"github.com/Arif14377/koda-b6-golang/internal/models"
	"github.com/Arif14377/koda-b6-golang/internal/services"
	"github.com/Arif14377/koda-b6-golang/internal/views"
)

// TODO: buat flow start app
func StartAppController() {
	views.Landing()
	menu := views.Input("Input menu: ")

	switch menu {
	case 1:
		LoginController()
	case 2:
		MenuProductController()
	}
}

func MenuProductController() {
	views.MenuProduct()
	menu := views.Input("Input menu: ")

	switch menu {
	case 1:
		ShowProductController()
	case 2:
		SearchProductController()
	case 3:
		FilterProductController()
	case 0:
		StartAppController()
	}
}

func ShowProductController() {
	var url string = "https://raw.githubusercontent.com/Arif14377/koda-b6-weekly1/refs/heads/main/data.json"
	var cart models.Cart

	views.Product(url)
	// TODO: Keranjang masih reset setiap masuk keranjang
	for true {
		fmt.Println("# What do you want:")
		fmt.Println("1. Masukkan ke keranjang")
		fmt.Println("2. Checkout")
		pilihan := views.Input("Masukkan angka pilihan: ")
		// var cart []models.CartItem

		// 1. keranjang; 2. checkout.
		switch pilihan {
		case 1:
			id := views.Input("Masukkan id product: ")
			qty := views.Input("Masukkan jumlah produk: ")

			cart.Cart = append(cart.Cart, services.AddCart(id, qty, url))
			fmt.Printf("Cart Anda: %v\n", cart.Cart)

			totalPrice := services.TotalPrice(cart.Cart)
			fmt.Printf("Total Belanja: Rp. %d\n", totalPrice)

			continue
		}
	}
}

func SearchProductController() {

}

func FilterProductController() {

}
