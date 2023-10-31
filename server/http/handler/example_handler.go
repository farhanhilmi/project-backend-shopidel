package handler

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
	"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
	exampleUsecase usecase.ExampleUsecase
}

func NewExampleHandler(exampleUsecase usecase.ExampleUsecase) *ExampleHandler {
	return &ExampleHandler{
		exampleUsecase: exampleUsecase,
	}
}

func (h *ExampleHandler) Example(c *gin.Context) {

}
