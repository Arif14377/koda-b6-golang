package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Arif14377/koda-b6-golang/internal/models"
)

func FetchProduct(url string) ([]models.Product, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("resp: %v", resp)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("body: %v", body)

	var products []models.Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("err unmarshal: %v", err)

	return products, nil
}
