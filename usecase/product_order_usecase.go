package usecase

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/shopspring/decimal"
)

type ProductOrderUsecase interface {
	CancelOrderBySeller(ctx context.Context, req dtousecase.ProductOrderRequest) (*dtousecase.ProductOrderResponse, error)
	ProcessedOrder(ctx context.Context, req dtousecase.ProductOrderRequest) (*dtousecase.ProductOrderResponse, error)
	CheckoutOrder(ctx context.Context, req dtousecase.CheckoutOrderRequest) (*dtousecase.CheckoutOrderResponse, error)
	CheckDeliveryFee(ctx context.Context, req dtousecase.CheckDeliveryFeeRequest) (*dtousecase.CourierFeeResponse, error)
	GetCouriers(ctx context.Context, req dtousecase.SellerCourier) ([]model.Couriers, error)
	GetOrderHistories(ctx context.Context, req dtousecase.ProductOrderHistoryRequest) ([]dtousecase.OrdersResponse, dtogeneral.PaginationData, error)
}

type productOrderUsecase struct {
	productOrderRepository    repository.ProductOrdersRepository
	productCombinationVariant repository.ProductVariantCombinationRepository
	accountAddressRepository  repository.AccountAddressRepository
	accountRepository         repository.AccountRepository
	courierRepository         repository.CourierRepository
}

type ProductOrderUsecaseConfig struct {
	ProductOrderRepository              repository.ProductOrdersRepository
	ProductVariantCombinationRepository repository.ProductVariantCombinationRepository
	AccountAddressRepository            repository.AccountAddressRepository
	AccountRepository                   repository.AccountRepository
	CourierRepository                   repository.CourierRepository
}

func NewProductOrderUsecase(config ProductOrderUsecaseConfig) ProductOrderUsecase {
	au := &productOrderUsecase{}
	if config.ProductOrderRepository != nil {
		au.productOrderRepository = config.ProductOrderRepository
	}
	if config.ProductVariantCombinationRepository != nil {
		au.productCombinationVariant = config.ProductVariantCombinationRepository
	}
	if config.AccountAddressRepository != nil {
		au.accountAddressRepository = config.AccountAddressRepository
	}
	if config.AccountRepository != nil {
		au.accountRepository = config.AccountRepository
	}
	if config.CourierRepository != nil {
		au.courierRepository = config.CourierRepository
	}

	return au
}

func (u *productOrderUsecase) CheckoutOrder(ctx context.Context, req dtousecase.CheckoutOrderRequest) (*dtousecase.CheckoutOrderResponse, error) {
	id, err := strconv.Atoi(req.DestinationAddressID)
	if err != nil {
		return nil, err
	}
	address, err := u.accountAddressRepository.FindBuyerAddressByID(ctx, dtorepository.AccountAddressRequest{ID: id})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	orderDetails := []dtorepository.ProductOrderDetailRequest{}
	totalPayment := decimal.NewFromInt(0)

	for _, p := range req.ProductVariant {
		if p.Quantity < 1 {
			return nil, util.ErrQtyInputZero
		}

		productVariant, err := u.productCombinationVariant.FindById(ctx, dtorepository.ProductCombinationVariantRequest{ID: p.ID})
		if errors.Is(err, util.ErrNoRecordFound) {
			return nil, util.ErrNoRecordFound
		}
		if err != nil {
			return nil, err
		}
		if productVariant.Stock < 1 {
			return nil, util.ErrInsufficientStock
		}

		order := dtorepository.ProductOrderDetailRequest{
			Quantity:                             p.Quantity,
			ProductVariantSelectionCombinationID: p.ID,
			IndividualPrice:                      productVariant.IndividualPrice,
		}
		qty, err := decimal.NewFromString(fmt.Sprintf("%v", p.Quantity))
		if err != nil {
			return nil, err
		}
		totalPayment = totalPayment.Add(productVariant.IndividualPrice.Mul(qty))
		orderDetails = append(orderDetails, order)
	}

	courier, err := u.courierRepository.FindById(ctx, dtorepository.CourierData{ID: req.CourierID})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	courierFee, err := u.CheckDeliveryFee(ctx, dtousecase.CheckDeliveryFeeRequest{
		ID:          req.CourierID,
		Destination: req.DestinationAddressID,
		Weight:      req.Weight,
		Courier:     courier.Name,
		SellerID:    req.SellerID,
	})
	if err != nil {
		return nil, err
	}

	deliveryFee := decimal.NewFromInt(int64(courierFee.Cost))
	totalPayment = totalPayment.Add(deliveryFee)

	account, err := u.accountRepository.FindById(ctx, dtorepository.GetAccountRequest{UserId: req.UserID})
	if err != nil {
		return nil, err
	}

	if account.WalletPin == "" {
		return nil, util.ErrWalletNotSet
	}

	if account.Balance.LessThan(totalPayment) {
		return nil, util.ErrInsufficientBalance
	}

	seller, err := u.accountRepository.FindById(ctx, dtorepository.GetAccountRequest{
		UserId: req.SellerID,
	})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	orderRequest := dtorepository.ProductOrderRequest{
		Province:           address.Province,
		District:           address.District,
		SubDistrict:        address.SubDistrict,
		Kelurahan:          address.Kelurahan,
		AddressDetail:      address.Detail,
		ZipCode:            address.ZipCode,
		AccountID:          req.UserID,
		SellerID:           req.SellerID,
		CourierID:          req.CourierID,
		Status:             constant.StatusOrderOnProcess,
		Notes:              req.Notes,
		DeliveryFee:        deliveryFee,
		TotalAmount:        totalPayment,
		TotalSellerAmount:  totalPayment.Sub(deliveryFee),
		ProductVariants:    orderDetails,
		BuyerWalletNumber:  account.WalletNumber,
		SellerWalletNumber: seller.WalletNumber,
	}

	order, err := u.productOrderRepository.Create(ctx, orderRequest)
	if err != nil {
		return nil, err
	}

	return &dtousecase.CheckoutOrderResponse{
		ID: order.ID,
	}, nil
}

func (u *productOrderUsecase) CancelOrderBySeller(ctx context.Context, req dtousecase.ProductOrderRequest) (*dtousecase.ProductOrderResponse, error) {
	order, err := u.productOrderRepository.FindByIDAndSellerID(ctx, dtorepository.ProductOrderRequest{
		ID:       req.ID,
		SellerID: req.SellerID,
		Status:   constant.StatusOrderOnProcess,
	})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	if len(order) < 1 {
		return nil, util.ErrOrderNotFound
	}

	increseProductsStock := []model.ProductCombinationVariant{}
	refundedAmount := decimal.NewFromInt(0)

	for _, v := range order {
		product := model.ProductCombinationVariant{}
		product.ID = v.ProductVariantSelectionCombinationID
		product.Stock = v.ProductStock + v.Quantity

		qty, err := decimal.NewFromString(fmt.Sprintf("%v", v.Quantity))
		if err != nil {
			return nil, err
		}

		increseProductsStock = append(increseProductsStock, product)
		refundedAmount = refundedAmount.Add(v.IndividualPrice.Mul(qty))
	}

	data, err := u.productOrderRepository.UpdateOrderStatusByIDAndAccountID(ctx, dtorepository.ReceiveOrderRequest{
		ID:                 req.ID,
		SellerID:           req.SellerID,
		Status:             constant.StatusCanceled,
		Notes:              req.Notes,
		TotalAmount:        refundedAmount,
		AccountID:          order[0].AccountID,
		Products:           increseProductsStock,
		SellerWalletNumber: req.SellerWalletNumber,
	})

	if err != nil {
		return nil, err
	}

	return &dtousecase.ProductOrderResponse{
		ID:     data.ID,
		Status: data.Status,
		Notes:  req.Notes,
	}, nil
}

func (u *productOrderUsecase) ProcessedOrder(ctx context.Context, req dtousecase.ProductOrderRequest) (*dtousecase.ProductOrderResponse, error) {
	order, err := u.productOrderRepository.FindByIDAndSellerID(ctx, dtorepository.ProductOrderRequest{
		ID:       req.ID,
		SellerID: req.SellerID,
		Status:   constant.StatusOrderOnProcess,
	})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	if len(order) < 1 {
		return nil, util.ErrOrderNotFound
	}

	totalAmount := decimal.NewFromInt(0)
	for _, v := range order {
		totalAmount.Add(v.IndividualPrice)
	}

	data, err := u.productOrderRepository.ProcessedOrder(ctx, dtorepository.ProductOrderRequest{
		ID:          req.ID,
		SellerID:    req.SellerID,
		Status:      constant.StatusProcessedOrder,
		TotalAmount: totalAmount,
	})

	if err != nil {
		return nil, err
	}

	return &dtousecase.ProductOrderResponse{
		ID:     data.ID,
		Status: data.Status,
	}, nil
}

func (u *productOrderUsecase) CheckDeliveryFee(ctx context.Context, req dtousecase.CheckDeliveryFeeRequest) (*dtousecase.CourierFeeResponse, error) {
	courier, err := u.courierRepository.FindById(ctx, dtorepository.CourierData{ID: req.ID})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	sellerAccount, err := u.accountAddressRepository.FindSellerAddressByAccountID(ctx, dtorepository.AccountAddressRequest{
		AccountID: req.SellerID,
	})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	req.Courier = courier.Name
	req.Origin = sellerAccount.RajaOngkirDistrictId
	response, err := util.GetRajaOngkirCost(req)
	if err != nil {
		return nil, err
	}

	courierRes := dtousecase.CourierFeeResponse{}

	for _, r := range response {
		for _, c := range r.Costs {
			if c.Service == "REG" {
				courierRes = dtousecase.CourierFeeResponse{
					Cost:      c.Cost[0].Value,
					Estimated: c.Cost[0].Etd,
					Note:      c.Cost[0].Note,
				}
			}
		}
	}

	return &courierRes, nil
}

func (u *productOrderUsecase) GetCouriers(ctx context.Context, req dtousecase.SellerCourier) ([]model.Couriers, error) {
	response, err := u.courierRepository.FindAllByShop(ctx, dtorepository.SellerCourier{SellerID: req.SellerID})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *productOrderUsecase) convertOrderHistoriesReponse(ctx context.Context, pagination dtogeneral.PaginationData, orders []model.ProductOrderHistories) ([]dtousecase.OrdersResponse, dtogeneral.PaginationData, error) {
	orderHistories := make(map[int][]dtousecase.OrderProduct)
	productOrders := []dtousecase.OrdersResponse{}

	for _, o := range orders {
		review := dtousecase.ProductOrderReview{}
		product := dtousecase.OrderProduct{
			ProductName:     o.ProductName,
			Quantity:        o.Quantity,
			IndividualPrice: o.IndividualPrice,
			ProductID:       o.ProductID,
		}
		if o.ReviewID > 0 {
			review.ReviewID = o.ReviewID
			review.ReviewFeedback = o.Feedback
			review.ReviewRating = o.Rating
			review.CreatedAt = o.ReviewCreatedAt
			product.IsReviewed = true
		}

		product.Review = review
		orderHistories[o.ID] = append(orderHistories[o.ID], product)
	}

	for k, products := range orderHistories {
		var totalAmount decimal.Decimal
		for _, o := range products {
			qty, err := decimal.NewFromString(fmt.Sprintf("%v", o.Quantity))
			if err != nil {
				return nil, dtogeneral.PaginationData{}, err
			}

			totalAmount = totalAmount.Add(o.IndividualPrice.Mul(qty))
		}
		order := dtousecase.OrdersResponse{
			OrderID:      k,
			Products:     products,
			TotalPayment: totalAmount,
		}
		productOrders = append(productOrders, order)
	}

	return productOrders, pagination, nil
}

func (u *productOrderUsecase) GetOrderHistories(ctx context.Context, req dtousecase.ProductOrderHistoryRequest) ([]dtousecase.OrdersResponse, dtogeneral.PaginationData, error) {
	var err error
	if req.Status == "" || strings.EqualFold(req.Status, constant.StatusOrderAll) {
		orders, totalItems, err := u.productOrderRepository.FindAllOrderHistoriesByUser(ctx, dtorepository.ProductOrderHistoryRequest{
			AccountID: req.AccountID,
			SortBy:    req.SortBy,
			Sort:      req.Sort,
			Limit:     req.Limit,
			Page:      req.Page,
			StartDate: req.StartDate,
			EndDate:   req.EndDate,
		})
		if err != nil {
			return nil, dtogeneral.PaginationData{}, err
		}

		pagination := dtogeneral.PaginationData{
			TotalItem:   int(totalItems),
			TotalPage:   (int(totalItems) + req.Limit - 1) / req.Limit,
			CurrentPage: req.Page,
			Limit:       req.Limit,
		}

		return u.convertOrderHistoriesReponse(ctx, pagination, orders)
	}

	orders, totalItems, err := u.productOrderRepository.FindAllOrderHistoriesByUserAndStatus(ctx, dtorepository.ProductOrderHistoryRequest{
		AccountID: req.AccountID,
		Status:    req.Status,
		SortBy:    req.SortBy,
		Sort:      req.Sort,
		Limit:     req.Limit,
		Page:      req.Page,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})
	if err != nil {
		return nil, dtogeneral.PaginationData{}, err
	}

	pagination := dtogeneral.PaginationData{
		TotalItem:   int(totalItems),
		TotalPage:   (int(totalItems) + req.Limit - 1) / req.Limit,
		CurrentPage: req.Page,
		Limit:       req.Limit,
	}

	return u.convertOrderHistoriesReponse(ctx, pagination, orders)
}
