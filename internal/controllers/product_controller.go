package controllers

import (
	"fmt"

	"github.com/Arif14377/koda-b6-golang/internal/views"
)

// TODO: buat flow start app
func StartAppController() {
	views.Landing()
	menu := views.Input("Input menu: ")
	fmt.Println(menu)

	switch menu {
	case 1:
		LoginController()
	case 2:
		ShowProductController()
	}
}

func ShowProductController() {
	views.Product()
}
