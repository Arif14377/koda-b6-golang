package controllers

import (
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

}

func SearchProductController() {

}

func FilterProductController() {

}
