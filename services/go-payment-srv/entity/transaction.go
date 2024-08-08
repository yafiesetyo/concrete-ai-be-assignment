package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Id                uint64          `json:"id"`
	Type              string          `json:"type"`
	Amount            decimal.Decimal `json:"amount"`
	Currency          string          `json:"currency"`
	Description       *string         `json:"description"`
	FromAccountNumber *string         `json:"from_account_number"`
	ToAccountNumber   *string         `json:"to_account_number"`
	Status            string          `json:"status"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
