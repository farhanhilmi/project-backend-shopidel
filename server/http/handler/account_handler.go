package handler

import (
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtohttp "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/http"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
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

func (h *AccountHandler) GetProfile(c *gin.Context) {
	var req dtohttp.GetAccountRequest
	req.UserId = c.GetInt("userId")
	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.GetAccountRequest{
		UserId: req.UserId,
	}

	uRes, err := h.accountUsecase.GetProfile(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtohttp.GetAccountResponse{
		ID:             uRes.ID,
		FullName:       uRes.FullName,
		Username:       uRes.Username,
		Email:          uRes.Email,
		PhoneNumber:    uRes.PhoneNumber,
		Gender:         uRes.Gender,
		Birthdate:      uRes.Birthdate,
		ProfilePicture: uRes.ProfilePicture,
		WalletNumber:   uRes.WalletNumber,
		Balance:        uRes.Balance,
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "successfully get profile detail", Data: res})
}

func (h *AccountHandler) ActivateMyWallet(c *gin.Context) {
	var payload dto.ActivateWalletRequest

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	userId := c.GetInt("userId")
	uReq := dtousecase.GetAccountRequest{
		UserId: userId,
	}

	payload = dto.ActivateWalletRequest{
		PIN: strings.TrimSpace(payload.PIN),
	}

	_, err = h.accountUsecase.ActivateMyWallet(c.Request.Context(), uReq, payload.PIN)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "successfully setup PIN"})
}

func (h *AccountHandler) ChangeWalletPIN(c *gin.Context) {
	var payload dtohttp.ChangeWalletPINRequest

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.UpdateWalletPINRequest{
		UserID:       c.GetInt("userId"),
		WalletNewPIN: strings.TrimSpace(payload.WalletNewPIN),
		WalletPIN:    strings.TrimSpace(payload.WalletPIN),
	}

	_, err = h.accountUsecase.ChangeMyWalletPIN(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "successfully update wallet PIN"})
}

func (h *AccountHandler) CheckISPasswordCorrect(c *gin.Context) {
	var payload dtohttp.CheckPasswordRequest

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.AccountRequest{
		ID:       c.GetInt("userId"),
		Password: strings.TrimSpace(payload.Password),
	}

	result, err := h.accountUsecase.CheckPasswordCorrect(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: result})
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req dtohttp.CreateAccountRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	req.Email = strings.TrimSpace(req.Email)
	req.FullName = strings.TrimSpace(req.FullName)
	req.Username = strings.TrimSpace(req.Username)

	uReq := dtousecase.CreateAccountRequest{
		Username: strings.TrimSpace(req.Username),
		FullName: strings.TrimSpace(req.FullName),
		Email:    strings.TrimSpace(req.Email),
		Password: req.Password,
	}

	uRes, err := h.accountUsecase.CreateAccount(c.Request.Context(), uReq)

	if err != nil {
		c.Error(err)
		return
	}

	res := dtohttp.CreateAccountResponse{
		Username: uRes.Username,
		FullName: uRes.FullName,
		Email:    uRes.Email,
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: res})
}
