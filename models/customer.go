package models

import "time"

type Customer struct {
	PurchaseDate time.Time
	Email        string
	CustomerID   string
}
