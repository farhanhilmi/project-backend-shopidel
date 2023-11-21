package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtohttp "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/http"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

type AccountHandler struct {
	accountUsecase             usecase.AccountUsecase
	myWalletTransactionUsecase usecase.MyWalletTransactionUsecase
}

func NewAccountHandler(accountUsecase usecase.AccountUsecase, myWalletTransactionUsecase usecase.MyWalletTransactionUsecase) *AccountHandler {
	return &AccountHandler{
		accountUsecase:             accountUsecase,
		myWalletTransactionUsecase: myWalletTransactionUsecase,
	}
}

func (h *AccountHandler) ChangePassword(c *gin.Context) {
	var req dtohttp.ChangePasswordRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.ChangePasswordRequest {
		AccountID:  c.GetInt("userId"),
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
	err = h.accountUsecase.ChangePassword(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}
	
	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Pasword Successfully Changed"})
}

func (h *AccountHandler) RequestChangePasswordOTP(c *gin.Context) {

	account, err := h.accountUsecase.RequestOTP(c.Request.Context(), dtousecase.ChangePasswordRequest{
		AccountID: c.GetInt("userId"),
	})
	if err != nil {
		c.Error(err)
		return
	}
	
	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: fmt.Sprintf("OTP Successfully sent to %s", account.Email)})
}

func (h *AccountHandler) RequestForgetPassword(c *gin.Context) {
	var req dtohttp.ForgetPasswordRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.ForgetPasswordRequest{
		Email: req.Email,
	}

	_, err = h.accountUsecase.RequestForgetPassword(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Forgot password request was successful. Please check your email to proceed"})
}

func (h *AccountHandler) RequestForgetChangePassword(c *gin.Context) {
	var req dtohttp.ForgetChangePasswordRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.ForgetChangePasswordRequest{
		Password: req.Password,
		Token:    req.Token,
	}

	_, err = h.accountUsecase.RequestForgetChangePassword(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Successfully set new password"})
}

func (h *AccountHandler) RegisterSeller(c *gin.Context) {
	res := dtohttp.RegisterSellerResponse{}
	var req dtohttp.RegisterSellerRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.RegisterSellerRequest{
		UserId:        c.GetInt("userId"),
		ShopName:      req.ShopName,
		AddressId:     req.AddressId,
		ListCourierId: req.ListCourierId,
	}

	uRes, err := h.accountUsecase.RegisterSeller(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res.ShopName = uRes.ShopName

	convertMessage := fmt.Sprintf("Merchant %s registered successfully", res.ShopName)
	c.JSON(http.StatusCreated, dtogeneral.JSONResponse{Message: convertMessage})
}

func (h *AccountHandler) GetAddresses(c *gin.Context) {
	res := []dtohttp.AddressResponse{}
	var req dtohttp.AddressRequest
	req.UserId = c.GetInt("userId")
	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.AddressRequest{
		UserId: req.UserId,
	}

	uRes, err := h.accountUsecase.GetAddresses(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	for _, data := range *uRes {
		res = append(res, dtohttp.AddressResponse{
			ID:              data.ID,
			FullAddress:     data.FullAddress,
			Detail:          data.Detail,
			ZipCode:         data.ZipCode,
			Kelurahan:       data.Kelurahan,
			SubDistrict:     data.SubDistrict,
			DistrictId:      data.DistrictId,
			District:        data.District,
			ProvinceId:      data.ProvinceId,
			Province:        data.Province,
			IsBuyerDefault:  data.IsBuyerDefault,
			IsSellerDefault: data.IsSellerDefault,
		})
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "successfully get addresses", Data: res})
}

func (h *AccountHandler) RegisterAdress(c *gin.Context) {
	req := dtohttp.RegisterAddressRequest{}
	res := dtohttp.RegisterAddressRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.RegisterAddressRequest{
		ProvinceId:  req.ProvinceId,
		DistrictId:  req.DistrictId,
		SubDistrict: req.SubDistrict,
		Kelurahan:   req.Kelurahan,
		ZipCode:     req.ZipCode,
		Detail:      req.Detail,
		AccountId:   c.GetInt("userId"),
	}

	uRes, err := h.accountUsecase.RegisterAccountAddress(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res.Detail = uRes.Detail
	res.DistrictId = uRes.DistrictId
	res.Kelurahan = uRes.Kelurahan
	res.ProvinceId = uRes.ProvinceId
	res.SubDistrict = uRes.SubDistrict
	res.ZipCode = uRes.ZipCode

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: res})
}

func (h *AccountHandler) UpdateAddress(c *gin.Context) {
	req := dtohttp.UpdateAddressRequest{}
	res := dtohttp.UpdateAddressRequest{}

	id := c.Param("addressId")
	addressId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.UpdateAddressRequest{
		AddressId:       addressId,
		ProvinceId:      req.ProvinceId,
		DistrictId:      req.DistrictId,
		SubDistrict:     req.SubDistrict,
		Kelurahan:       req.Kelurahan,
		ZipCode:         req.ZipCode,
		Detail:          req.Detail,
		IsBuyerDefault:  *req.IsBuyerDefault,
		IsSellerDefault: *req.IsSellerDefault,
		AccountId:       c.GetInt("userId"),
	}

	uRes, err := h.accountUsecase.UpdateAccountAddress(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res.Detail = uRes.Detail
	res.DistrictId = uRes.DistrictId
	res.Kelurahan = uRes.Kelurahan
	res.ProvinceId = uRes.ProvinceId
	res.SubDistrict = uRes.SubDistrict
	res.ZipCode = uRes.ZipCode
	res.IsBuyerDefault = &uRes.IsBuyerDefault
	res.IsSellerDefault = &uRes.IsSellerDefault

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: res})
}

func (h *AccountHandler) DeleteAdress(c *gin.Context) {
	addressIdString := c.Param("addressId")
	addressId, err := strconv.Atoi(addressIdString)
	if err != nil {
		c.Error(err)
		return
	}

	err = h.accountUsecase.DeleteAddresses(c.Request.Context(), dtousecase.DeleteAddressRequest{AddressId: addressId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "success deleting the address"})
}

func (h *AccountHandler) Login(c *gin.Context) {
	var req dtohttp.LoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	uRes, err := h.accountUsecase.Login(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtogeneral.JSONResponse{
		AccessToken:  uRes.AccessToken,
		RefreshToken: uRes.RefreshToken,
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Login success", Data: res})
}

func (h *AccountHandler) RefreshToken(c *gin.Context) {
	var req dtohttp.RefreshTokenRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	}

	uRes, err := h.accountUsecase.RefreshToken(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	res := dtogeneral.JSONResponse{
		AccessToken:  uRes.AccessToken,
		RefreshToken: uRes.RefreshToken,
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: res})
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
		IsSeller:       uRes.IsSeller,
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

func (h *AccountHandler) ValidateWalletPIN(c *gin.Context) {
	var payload dtohttp.ValidateWAlletPINRequest

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.ValidateWAlletPINRequest{
		UserID:    c.GetInt("userId"),
		WalletPIN: strings.TrimSpace(payload.WalletPIN),
	}

	result, err := h.accountUsecase.ValidateWalletPIN(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: dtohttp.ValidateWAlletPINResponse{IsCorrect: result.IsCorrect}})
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

func (h *AccountHandler) GetCart(c *gin.Context) {
	uReq := dtousecase.GetCartRequest{
		UserId: c.GetInt("userId"),
	}

	uRes, err := h.accountUsecase.GetCart(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: uRes})
}

func (h *AccountHandler) AddProductToCart(c *gin.Context) {
	req := dtohttp.AddProductToCartRequest{}
	res := dtohttp.AddProductToCartResponse{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	uReq := dtousecase.AddProductToCartRequest{
		UserId:           c.GetInt("userId"),
		ProductVariantId: req.ProductId,
		Quantity:         req.Quantity,
	}

	uRes, err := h.accountUsecase.AddProductToCart(c.Request.Context(), uReq)
	res.ProductId = uRes.ProductId
	res.Quantity = uRes.Quantity

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: res})
}

func (h *AccountHandler) GetListTransactions(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.Error(err)
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.Error(err)
		return
	}
	sortBy := c.DefaultQuery("sortBy", "date")
	sort := c.DefaultQuery("sort", "desc")
	startDate := c.DefaultQuery("startDate", "")
	endDate := c.DefaultQuery("endDate", "")
	transactionType := c.DefaultQuery("type", "")

	if valid := util.IsDateValid(startDate); !valid && startDate != "" {
		c.Error(util.ErrInvalidDateFormat)
		return
	}
	if valid := util.IsDateValid(endDate); !valid && endDate != "" {
		c.Error(util.ErrInvalidDateFormat)
		return
	}

	if !slices.Contains([]string{"date", "amount", "type"}, sortBy) {
		c.Error(util.ErrWalletHistorySortBy)
		return
	}

	switch sortBy {
	case "date":
		sortBy = "created_at"
	}

	reqWallet := dtousecase.WalletHistoriesParams{
		SortBy:          sortBy,
		Sort:            sort,
		Limit:           limit,
		Page:            page,
		StartDate:       startDate,
		EndDate:         endDate,
		AccountID:       c.GetInt("userId"),
		TransactionType: transactionType,
	}

	transactions, pagination, err := h.myWalletTransactionUsecase.GetTransactions(c.Request.Context(), reqWallet)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONPagination{Data: transactions, Pagination: *pagination})
}

func (h *AccountHandler) UpdateCart(c *gin.Context) {
	var payload dtohttp.UpdateCartRequest

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.UpdateCartRequest{
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
	}

	uRes, err := h.accountUsecase.UpdateCartQuantity(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: uRes})
}

func (h *AccountHandler) DeleteCartProduct(c *gin.Context) {
	var payload dtohttp.DeleteCartProductRequest

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.Error(util.ErrInvalidInput)
		return
	}

	uReq := dtousecase.DeleteCartProductRequest{
		ListProductID: payload.ListProductID,
	}

	_, err = h.accountUsecase.DeleteProductCart(c.Request.Context(), uReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Message: "Successfully delete cart"})
}

func (h *AccountHandler) GetCouriers(c *gin.Context) {
	response, err := h.accountUsecase.GetCouriers(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}

func (h *AccountHandler) GetProvinces(c *gin.Context) {
	response, err := h.accountUsecase.GetProvinces(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}

func (h *AccountHandler) GetDistricts(c *gin.Context) {
	response, err := h.accountUsecase.GetDistricts(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response.Districts})
}

func (h *AccountHandler) GetProvinceDistricts(c *gin.Context) {
	provinceIdString := c.Param("provinceId")

	provinceId, err := strconv.Atoi(provinceIdString)
	if err != nil {
		c.Error(err)
		return
	}

	response, err := h.accountUsecase.GetDistrictsByProvinceId(c.Request.Context(), dtousecase.GetDistrictRequest{ProvinceId: provinceId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}

func (h *AccountHandler) GetCategories(c *gin.Context) {
	response, err := h.accountUsecase.GetCategories(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtogeneral.JSONResponse{Data: response})
}
