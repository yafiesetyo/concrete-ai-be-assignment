package handler

import (
	"encoding/json"
	"go-payment-srv/request"
	validationerror "go-payment-srv/validationError"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Send(c *gin.Context) {
	var req request.Send

	email := c.MustGet("email").(string)
	req.Email = email

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := h.usecase.Send(c.Request.Context(), req)
	if h.isClientErr(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}

func (h *handler) Withdraw(c *gin.Context) {
	var req request.Withdraw

	email := c.MustGet("email").(string)
	req.Email = email

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := h.usecase.Withdraw(c.Request.Context(), req)
	if h.isClientErr(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}

func (h *handler) Auth(c *gin.Context) {
	authorization := c.Request.Header["Authorization"]

	if len(authorization) < 1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "authorization not provided",
		})
		return
	}

	token := authorization[0]
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "token not provided",
		})
		return
	}

	req, err := http.NewRequest(http.MethodGet, h.nodeAccountUrl+"/accounts/session", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := h.httpClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	bt, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	type responseStruct struct {
		Message string `json:"message"`
		Data    struct {
			Email string `json:"email"`
		} `json:"data"`
	}
	response := responseStruct{}

	if err := json.Unmarshal(bt, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if resp.StatusCode != 200 {
		c.JSON(resp.StatusCode, gin.H{
			"message": response.Message,
		})
		return
	}

	c.Set("email", response.Data.Email)
	c.Next()
}

func (h *handler) isClientErr(err error) bool {
	if err == nil {
		return false
	}

	e := map[error]bool{
		validationerror.ErrAccountNotFound:         true,
		validationerror.ErrInsufficientBalance:     true,
		validationerror.ErrInvalidAccount:          true,
		validationerror.ErrReceiverAccountNotFound: true,
	}

	isOk := e[err]
	return isOk
}
