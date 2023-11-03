package usecase

import (
	"context"
	"errors"
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/shopspring/decimal"
)

type ProductOrderUsecase interface {
	CancelOrderBySeller(ctx context.Context, req dtousecase.ProductOrderRequest) (*dtousecase.ProductOrderResponse, error)
	ProcessedOrder(ctx context.Context, req dtousecase.ProductOrderRequest) (*dtousecase.ProductOrderResponse, error)
}

type productOrderUsecase struct {
	productOrderRepository repository.ProductOrdersRepository
}

type ProductOrderUsecaseConfig struct {
	ProductOrderRepository repository.ProductOrdersRepository
}

func NewProductOrderUsecase(config ProductOrderUsecaseConfig) ProductOrderUsecase {
	au := &productOrderUsecase{}
	if config.ProductOrderRepository != nil {
		au.productOrderRepository = config.ProductOrderRepository
	}

	return au
}

func (u *productOrderUsecase) CancelOrderBySeller(ctx context.Context, req dtousecase.ProductOrderRequest) (*dtousecase.ProductOrderResponse, error) {
	order, err := u.productOrderRepository.FindByIDAndSellerID(ctx, dtorepository.ProductOrderRequest{
		ID:       req.ID,
		SellerID: req.SellerID,
		Status:   constant.StatusWaitingSellerConfirmation,
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

	refundedAmount := decimal.NewFromInt(0)
	for _, v := range order {
		qty, err := decimal.NewFromString(fmt.Sprintf("%v", v.Quantity))
		if err != nil {
			return nil, err
		}
		refundedAmount = refundedAmount.Add(v.IndividualPrice.Mul(qty))
	}

	// TODO: add product stock when seller cancel the order

	data, err := u.productOrderRepository.UpdateOrderStatusByIDAndAccountID(ctx, dtorepository.ProductOrderRequest{
		ID:          req.ID,
		SellerID:    req.SellerID,
		Status:      constant.StatusCanceled,
		Notes:       req.Notes,
		TotalAmount: refundedAmount,
		AccountID:   order[0].AccountID,
	})

	if err != nil {
		return nil, err
	}

	return &dtousecase.ProductOrderResponse{
		ID:     data.ID,
		Status: data.Status,
		Notes:  data.Notes,
	}, nil
}

func (u *productOrderUsecase) ProcessedOrder(ctx context.Context, req dtousecase.ProductOrderRequest) (*dtousecase.ProductOrderResponse, error) {
	order, err := u.productOrderRepository.FindByIDAndSellerID(ctx, dtorepository.ProductOrderRequest{
		ID:       req.ID,
		SellerID: req.SellerID,
		Status:   constant.StatusWaitingSellerConfirmation,
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
