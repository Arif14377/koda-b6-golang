package services

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/Arif14377/koda-b6-golang/internal/models"
)

func SearchProduct(searchKeyword string, cart []models.Product) ([]models.Product, error) {
	if reflect.TypeOf(searchKeyword).String() != "string" {
		var strKosong []models.Product
		return strKosong, errors.New("error")
	}

	fmt.Println("searchKeyword = ", searchKeyword)
	var result []models.Product
	for x := range cart {
		if strings.Contains(cart[x].Name, searchKeyword) {
			result = append(result, cart[x])
		}
	}
	return result, nil
}
