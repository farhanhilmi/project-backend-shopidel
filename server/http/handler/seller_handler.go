package handler

import (
	"net/http"
	"strconv"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
	"github.com/gin-gonic/gin"
)

type SellerHandler struct {
	sellerUsecase usecase.SellerUsecase
}

type SellerHandlerConfig struct {
	SellerUsecase usecase.SellerUsecase
}

func NewSellerHandler(config SellerHandlerConfig) *SellerHandler {
	ah := &SellerHandler{}

	if config.SellerUsecase != nil {
		ah.sellerUsecase = config.SellerUsecase
	}

	return ah
}

func (h *SellerHandler) GetProfile(c *gin.Context) {
	sellerIdString := c.Param("sellerId")
	sellerId, err := strconv.Atoi(sellerIdString)
	if err != nil {
		c.Error(err)
		return
	}

	uRes, err := h.sellerUsecase.GetProfile(c.Request.Context(), dtousecase.GetSellerProfileRequest{SellerId: sellerId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "successfully get profile detail", Data: uRes})
}
