package model

import (
	"go-payment-srv/entity"
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Id                uint64          `gorm:"id"`
	Type              string          `gorm:"column:type"`
	Amount            decimal.Decimal `gorm:"column:amount"`
	Currency          string          `gorm:"column:currency"`
	Description       *string         `gorm:"column:description"`
	FromAccountNumber *string         `gorm:"column:from_account_number"`
	ToAccountNumber   *string         `gorm:"column:to_account_number"`
	Status            string          `gorm:"column:status"`

	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (t *Transaction) FromEntity(in entity.Transaction) {
	if t == nil {
		return
	}

	t.Id = in.Id
	t.Type = in.Type
	t.Amount = in.Amount
	t.Currency = in.Currency
	t.Description = in.Description
	t.FromAccountNumber = in.FromAccountNumber
	t.ToAccountNumber = in.ToAccountNumber
	t.Status = in.Status
}
