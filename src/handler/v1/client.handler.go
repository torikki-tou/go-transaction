package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/torikki-tou/go-transaction/common"
	"github.com/torikki-tou/go-transaction/dto"
	"github.com/torikki-tou/go-transaction/service"
)

type ClientHandler interface {
	ChangeBalance(ctx *gin.Context)
}

type clientHandler struct {
	clientService service.ClientService
}

func NewClientHandler(clientService service.ClientService) ClientHandler {
	return &clientHandler{
		clientService: clientService,
	}
}

func (c clientHandler) ChangeBalance(ctx *gin.Context) {
	var changeBalanceRequest dto.ChangeBalance
	err := ctx.ShouldBind(&changeBalanceRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
	}

	err = c.clientService.ChangeBalance(changeBalanceRequest)
	if err != nil {
		if errors.Is(err, &common.LowBalanceError{}) {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		} else if errors.Is(err, &common.InternalBDError{}) {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		} else if errors.Is(err, &common.ClientNotFoundError{}) {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else if errors.Is(err, &common.NotificationError{}) {
			ctx.JSON(http.StatusOK, gin.H{"warning": err.Error()})
		}
	}

	ctx.Status(http.StatusOK)
}
