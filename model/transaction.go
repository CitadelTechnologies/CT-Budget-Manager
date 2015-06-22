package model

import "time"

type(
	Transaction struct {
		Wording string `json:"wording"`
		Description string `json:"description"`
		Type bool `json:"type"`
		Sector Sector `json:"sector"`
		Amount int `json:"amount"`
		Date time.Time `json:"date"`
	}
	Transactions []Transaction
)
