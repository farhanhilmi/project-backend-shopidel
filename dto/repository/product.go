package dtorepository

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"github.com/shopspring/decimal"
)

type ProductRequest struct {
	ProductID  int
	Name       string
	CategoryID int
	IsActive   bool
}

type ProductResponse struct {
	ID                int
	Name              string
	Description       string
	CategoryID        int
	Category          model.Category
	ProductImages     model.ProductImages
	ProductVideos     model.ProductVideos
	HazardousMaterial bool
	Weight            decimal.Decimal
	Size              decimal.Decimal
	IsNew             bool
	InternalSKU       string
	ViewCount         int
	IsActive          bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}
