package request

import "github.com/shopspring/decimal"

type Send struct {
	Email       string          `json:"-"`
	Number      string          `json:"number" binding:"required"`
	To          string          `json:"to" binding:"required"`
	Amount      decimal.Decimal `json:"amount" binding:"required"`
	Currency    string          `json:"currency" binding:"required"`
	Description *string         `json:"description"`
}
