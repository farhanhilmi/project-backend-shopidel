package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
)

type exampleUsecase struct {
	exampleRepo repository.ExampleRepository
}

type ExampleUsecase interface {
	Example(ctx context.Context) (*dto.Res, error)
}

func NewAuthenticationUsecase(exampleRepo repository.ExampleRepository) ExampleUsecase {
	return &exampleUsecase{}
}

func (u *exampleUsecase) Example(ctx context.Context) (*dto.Res, error) {

}
