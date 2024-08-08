package request

import "github.com/shopspring/decimal"

type Withdraw struct {
	Email    string          `json:"-"`
	Number   string          `json:"number" binding:"required"`
	Amount   decimal.Decimal `json:"amount" binding:"required"`
	Currency string          `json:"currency" binding:"required"`
}
