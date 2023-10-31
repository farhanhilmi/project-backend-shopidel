package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
)

type accountUsecase struct {
	accountRepo repository.AccountRepository
}

type AccountUsecase interface {
	GetDetail(ctx context.Context) (*dto.AccountResponse, error)
}

func NewAuthenticationUsecase(accountRepo repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		accountRepo: accountRepo,
	}
}

func (u *accountUsecase) GetDetail(ctx context.Context) (*dto.AccountResponse, error) {
	accountDetail, err := u.accountRepo.GetDetail(ctx)
	if err != nil {
		return &dto.AccountResponse{}, err
	}

	return accountDetail, nil
}
