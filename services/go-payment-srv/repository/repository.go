package repository

import (
	"go-payment-srv/entity"
	"go-payment-srv/repository/model"

	"gorm.io/gorm"
)

func (r *repo) Transaction(fc func(tx *gorm.DB) error) error {
	return r.db.Transaction(fc)
}

func (r *repo) CreateTransaction(tx *gorm.DB, transaction model.Transaction) error {
	return tx.
		Table(transaction.TableName()).
		Create(&transaction).
		Error
}

func (r *repo) UpdateTransaction(tx *gorm.DB, transaction model.Transaction) error {
	return tx.
		Table(transaction.TableName()).
		Where(`id=?`, transaction.Id).
		Updates(&transaction).
		Error
}

func (r *repo) UpdateAccount(tx *gorm.DB, account model.Account) error {
	return tx.
		Table(account.TableName()).
		Where(`id=?`, account.Id).
		Updates(&account).
		Error
}

func (r *repo) GetAccountByNumber(number string) (entity.Account, error) {
	var (
		account model.Account
	)

	err := r.db.
		Table(account.TableName()).
		Where(`"number"=?`, number).
		First(&account).Error
	if err != nil {
		return entity.Account{}, err
	}

	return account.ToEntity(), nil
}

func (r *repo) GetEmailByAccountNumber(number string) (string, error) {
	var (
		email string
	)

	return email, r.db.Raw(`
	select u.email from users u 
	join accounts a on a.user_id = u.id
	where a."number" = ?
	`, number).Scan(&email).Error
}
