package usecase

import (
	"context"
	"errors"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/constant"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
)

type ProductOrderUsecase interface {
	CancelOrderBySeller(ctx context.Context, req dtousecase.ProductOrderRequest) (*dtousecase.ProductOrderResponse, error)
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
	order, err := u.productOrderRepository.FindByID(ctx, dtorepository.ProductOrderRequest{
		ID: req.ID,
	})
	if errors.Is(err, util.ErrNoRecordFound) {
		return nil, util.ErrNoRecordFound
	}
	if err != nil {
		return nil, err
	}

	// TODO: check is the order belong to this seller

	if order.Status != constant.StatusWaitingSellerConfirmation {
		return nil, util.ErrOrderStatusNotWaiting
	}

	data, err := u.productOrderRepository.UpdateOrderStatusByIDAndAccountID(ctx, dtorepository.ProductOrderRequest{
		ID:       req.ID,
		SellerID: req.SellerID,
		Status:   constant.StatusCanceled,
		Notes:    req.Notes,
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
