package handler

import (
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountUsecase usecase.AccountUsecase
}

func NewAccountHandler(accountUsecase usecase.AccountUsecase) *AccountHandler {
	return &AccountHandler{
		accountUsecase: accountUsecase,
	}
}

func (h *AccountHandler) ActivateMyWallet(c *gin.Context) {
	var payload dto.ActivateWalletRequest

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	userId := c.GetInt("userId")

	payload = dto.ActivateWalletRequest{
		PIN: strings.TrimSpace(payload.PIN),
	}

	_, err = h.accountUsecase.ActivateMyWallet(c.Request.Context(), userId, payload.PIN)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.JSONResponse{Message: "successfully setup PIN"})

}
