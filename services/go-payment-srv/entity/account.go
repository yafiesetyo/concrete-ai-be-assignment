package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Account struct {
	Id      uint64          `json:"id"`
	Number  *string         `json:"number"`
	UserId  uint64          `json:"user_id"`
	Balance decimal.Decimal `json:"balance"`
	Type    string          `json:"type"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
