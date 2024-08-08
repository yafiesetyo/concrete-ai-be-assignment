package handler

import (
	"go-payment-srv/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Send(c *gin.Context)
	Withdraw(c *gin.Context)

	Auth(c *gin.Context)
}

type handler struct {
	usecase        usecase.IUsecase
	httpClient     *http.Client
	nodeAccountUrl string
}

var _ (IHandler) = (*handler)(nil)

func NewHandler(usecase usecase.IUsecase, nodeAccountUrl string) *handler {
	return &handler{
		usecase:        usecase,
		httpClient:     &http.Client{},
		nodeAccountUrl: nodeAccountUrl,
	}
}
