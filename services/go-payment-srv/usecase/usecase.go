package usecase

import (
	"context"
	"fmt"
	"go-payment-srv/entity"
	"go-payment-srv/repository/model"
	"go-payment-srv/request"
	validationerror "go-payment-srv/validationError"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (uc *usecase) Send(ctx context.Context, in request.Send) error {
	email, err := uc.repo.GetEmailByAccountNumber(in.Number)
	if err != nil {
		uc.logger.Info("repo.GetEmailByAccountNumber error", zap.Error(err))
		return err
	}
	if email != in.Email {
		return validationerror.ErrInvalidAccount
	}

	// check balance
	sender, err := uc.repo.GetAccountByNumber(in.Number)
	if err != nil {
		uc.logger.Info("repo.GetAccountByNumber error", zap.Error(err))
		return err
	}

	if sender.Id < 1 {
		return validationerror.ErrAccountNotFound
	}
	if in.Amount.GreaterThan(sender.Balance) {
		return validationerror.ErrInsufficientBalance
	}

	receiver, err := uc.repo.GetAccountByNumber(in.To)
	if err != nil {
		uc.logger.Info("[receiver] repo.GetAccountByNumber error", zap.Error(err))
		return err
	}

	if receiver.Id < 1 {
		return validationerror.ErrReceiverAccountNotFound
	}

	return uc.repo.Transaction(func(tx *gorm.DB) (err error) {
		// create transaction
		var outcome model.Transaction
		outcome.FromEntity(entity.Transaction{
			Type:              "OUTCOME",
			Amount:            in.Amount.Neg(),
			Currency:          in.Currency,
			Description:       in.Description,
			FromAccountNumber: &in.Number,
			ToAccountNumber:   &in.To,
			Status:            "SUCCESS",
		})

		err = uc.repo.CreateTransaction(tx, outcome)
		if err != nil {
			uc.logger.Error("failed to create transaction [OUT]", zap.Error(err))
			return
		}

		var income model.Transaction
		income.FromEntity(entity.Transaction{
			Type:              "INCOME",
			Amount:            in.Amount,
			Currency:          in.Currency,
			Description:       in.Description,
			FromAccountNumber: &in.Number,
			ToAccountNumber:   &in.To,
			Status:            "SUCCESS",
		})

		err = uc.repo.CreateTransaction(tx, income)
		if err != nil {
			uc.logger.Error("failed to create transaction [IN]", zap.Error(err))
			return
		}

		// update account (sender--)
		now := time.Now()
		sender.Balance = sender.Balance.Sub(in.Amount)
		sender.UpdatedAt = &now

		var senderModel model.Account
		senderModel.FromEntity(sender)

		err = uc.repo.UpdateAccount(tx, senderModel)
		if err != nil {
			uc.logger.Error("failed to update account [sender]", zap.Error(err))
			return
		}

		// update account (receiver++)
		now = time.Now()
		receiver.Balance = receiver.Balance.Add(in.Amount)
		receiver.UpdatedAt = &now

		var receiverModel model.Account
		receiverModel.FromEntity(receiver)

		err = uc.repo.UpdateAccount(tx, receiverModel)
		if err != nil {
			uc.logger.Error("failed to update account [receiver]", zap.Error(err))
			return
		}

		uc.logger.Info("transaction success")
		return
	})
}

func (uc *usecase) Withdraw(ctx context.Context, in request.Withdraw) error {
	email, err := uc.repo.GetEmailByAccountNumber(in.Number)
	if err != nil {
		uc.logger.Info("repo.GetEmailByAccountNumber error", zap.Error(err))
		return err
	}
	if in.Email != email {
		return validationerror.ErrInvalidAccount
	}

	// check balance
	withdrawer, err := uc.repo.GetAccountByNumber(in.Number)
	if err != nil {
		uc.logger.Info("repo.GetAccountByNumber error", zap.Error(err))
		return err
	}

	if withdrawer.Id < 1 {
		return validationerror.ErrAccountNotFound
	}
	if in.Amount.GreaterThan(withdrawer.Balance) {
		return validationerror.ErrInsufficientBalance
	}

	return uc.repo.Transaction(func(tx *gorm.DB) (err error) {
		var outcome model.Transaction
		description := fmt.Sprintf("%s withdraw", *withdrawer.Number)
		outcome.FromEntity(entity.Transaction{
			Type:              "OUTCOME",
			Amount:            in.Amount.Neg(),
			Currency:          in.Currency,
			Description:       &description,
			FromAccountNumber: &in.Number,
			ToAccountNumber:   nil,
			Status:            "SUCCESS",
		})

		err = uc.repo.CreateTransaction(tx, outcome)
		if err != nil {
			uc.logger.Error("failed to create transaction [OUT]", zap.Error(err))
			return
		}

		now := time.Now()
		withdrawer.Balance = withdrawer.Balance.Sub(in.Amount)
		withdrawer.UpdatedAt = &now

		var withdrawerModel model.Account
		withdrawerModel.FromEntity(withdrawer)

		err = uc.repo.UpdateAccount(tx, model.Account(withdrawer))
		if err != nil {
			uc.logger.Error("failed to update account [receiver]", zap.Error(err))
			return
		}

		uc.logger.Info("transaction success")
		return
	})
}
