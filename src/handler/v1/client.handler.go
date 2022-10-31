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
	userService service.UserService
}

func NewClientHandler(userService service.UserService) ClientHandler {
	return &clientHandler{
		userService: userService,
	}
}

func (c clientHandler) ChangeBalance(ctx *gin.Context) {
	var changeBalanceRequest dto.ChangeBalance
	err := ctx.ShouldBind(&changeBalanceRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
	}

	err = c.userService.ChangeBalance(changeBalanceRequest)
	if err != nil {
		if errors.Is(err, &common.LowBalanceError{}) {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		} else {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	ctx.Status(http.StatusOK)
}
