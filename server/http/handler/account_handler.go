package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
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

func (h *AccountHandler) GetDetail(c *gin.Context) {
	detailProfile, err := h.accountUsecase.GetDetail(c)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, dto.JSONResponse{Message: "success", Data: detailProfile})
}


