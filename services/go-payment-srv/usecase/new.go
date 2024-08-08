package usecase

import (
	"context"
	"go-payment-srv/repository"
	"go-payment-srv/request"

	"go.uber.org/zap"
)

type IUsecase interface {
	Send(ctx context.Context, in request.Send) error
	Withdraw(ctx context.Context, in request.Withdraw) error
}

type usecase struct {
	repo   repository.IRepo
	logger *zap.SugaredLogger
}

var _ (IUsecase) = (*usecase)(nil)

func NewUsecase(repo repository.IRepo, logger *zap.SugaredLogger) *usecase {
	return &usecase{
		repo:   repo,
		logger: logger,
	}
}
