package model

import (
	"go-payment-srv/entity"
	"time"

	"github.com/shopspring/decimal"
)

type Account struct {
	Id      uint64          `gorm:"column:id"`
	Number  *string         `gorm:"column:number"`
	UserId  uint64          `gorm:"column:user_id"`
	Balance decimal.Decimal `gorm:"column:balance"`
	Type    string          `gorm:"column:type"`

	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (Account) TableName() string {
	return "accounts"
}

func (a Account) ToEntity() entity.Account {
	return entity.Account{
		Id:        a.Id,
		Number:    a.Number,
		Balance:   a.Balance,
		Type:      a.Type,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func (a *Account) FromEntity(in entity.Account) {
	if a == nil {
		return
	}

	a.Id = in.Id
	a.Balance = in.Balance
	a.Number = in.Number
	a.Type = in.Type
	a.UpdatedAt = in.UpdatedAt
}
