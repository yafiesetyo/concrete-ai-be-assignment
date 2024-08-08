package repository

import (
	"go-payment-srv/entity"
	"go-payment-srv/repository/model"

	"gorm.io/gorm"
)

type IRepo interface {
	Transaction(fc func(tx *gorm.DB) error) error
	CreateTransaction(tx *gorm.DB, transaction model.Transaction) error
	UpdateTransaction(tx *gorm.DB, transaction model.Transaction) error
	UpdateAccount(tx *gorm.DB, account model.Account) error
	GetAccountByNumber(number string) (entity.Account, error)
	GetEmailByAccountNumber(number string) (string, error)
}

type repo struct {
	db *gorm.DB
}

var _ (IRepo) = (*repo)(nil)

func NewRepo(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}
