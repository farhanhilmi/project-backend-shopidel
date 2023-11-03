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
	"github.com/go-playground/validator/v10"
)

type AccountHandler struct {
	accountUsecase usecase.AccountUsecase
}

func NewAccountHandler(accountUsecase usecase.AccountUsecase) *AccountHandler {
	return &AccountHandler{
		accountUsecase: accountUsecase,
	}
}

func (h *AccountHandler) Login(c *gin.Context) {
	var req dtohttp.LoginRequest
	
	v := validator.New()
	if err := v.Struct(req); err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.LoginRequest {
		Email: req.Email,
		Password: req.Password,
	}

	uRes, err := h.accountUsecase.Login(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtogeneral.JSONResponse {
		AccessToken: uRes.AccessToken,
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Login success", Data: res})
}

func (h *AccountHandler) EditProfile(c *gin.Context) {
	var req dtohttp.EditAccountRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.EditAccountRequest{
		UserId:         c.GetInt("userId"),
		FullName:       req.FullName,
		Username:       req.Username,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		Gender:         req.Gender,
		Birthdate:      req.Birthdate,
		ProfilePicture: req.ProfilePicture,
	}

	uRes, err := h.accountUsecase.EditProfile(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtohttp.EditAccountResponse{
		ID:             uRes.ID,
		FullName:       uRes.FullName,
		Username:       uRes.Username,
		Email:          uRes.Email,
		PhoneNumber:    uRes.PhoneNumber,
		Gender:         uRes.Gender,
		Birthdate:      uRes.Birthdate,
		ProfilePicture: uRes.ProfilePicture,
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "successfully edited profile", Data: res})
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

func (h *AccountHandler) TopUpBalanceWallet(c *gin.Context) {
	var payload dtohttp.TopUpBalanceWalletRequest

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.TopUpBalanceWalletRequest{
		UserID: c.GetInt("userId"),
		Amount: payload.Amount,
	}

	_, err = h.accountUsecase.TopUpBalanceWallet(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "successfully top up wallet balance"})
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

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: dtohttp.CheckPasswordResponse{IsCorrect: result.IsCorrect}})
}

func (h *AccountHandler) GetWallet(c *gin.Context) {
	uReq := dtousecase.AccountRequest{
		ID: c.GetInt("userId"),
	}

	result, err := h.accountUsecase.GetWallet(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	resWallet := dtohttp.WalletResponse{
		Balance:      result.Balance,
		WalletNumber: result.WalletNumber,
		IsActive:     result.IsActive,
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: resWallet})
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req dtohttp.CreateAccountRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.CreateAccountRequest{
		Username: strings.TrimSpace(req.Username),
		FullName: strings.TrimSpace(req.FullName),
		Email:    strings.TrimSpace(req.Email),
		Password: strings.TrimSpace(req.Password),
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
