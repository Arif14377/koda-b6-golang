package models

import "time"

type Order struct {
	id    int
	Order Cart
	Date  time.Time
}
