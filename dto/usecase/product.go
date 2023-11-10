package dtousecase

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

type GetProductDetailRequest struct {
	ProductId int
}

type GetProductDetailResponse struct {
	Id             int              `json:"id"`
	ProductName    string           `json:"name"`
	Description    string           `json:"description"`
	Stars          decimal.Decimal  `json:"stars"`
	Sold           int              `json:"sold"`
	Available      int              `json:"available"`
	Images         []string         `json:"images"`
	VariantOptions []VariantOption  `json:"variant_options,omitempty"`
	Variants       []ProductVariant `json:"variants,omitempty"`
}

type VariantOption struct {
	VariantOptionName string   `json:"variant_option_name"`
	Childs            []string `json:"childs"`
}

type ProductVariant struct {
	VariantId   int                `json:"variant_id"`
	VariantName string             `json:"variant_name"`
	Selections  []ProductSelection `json:"selections,omitempty"`
	Stock       int                `json:"stock"`
	Price       decimal.Decimal    `json:"price"`
}

type ProductSelection struct {
	SelectionVariantName string `json:"selection_variant_name,omitempty"`
	SelectionName        string `json:"selection_name,omitempty"`
}

type UpdateCartRequest struct {
	ProductID int
	Quantity  int
}

type UpdateCartResponse struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ProductCart struct {
	ID        int
	ProductID int
	Quantity  int
	AccountID int
}

type FavoriteProduct struct {
	ID        int
	ProductID int
	AccountID int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ProductFavoritesParams struct {
	AccountID int
	SortBy    string
	Search    string
	Sort      string
	Limit     int
	Page      int
	StartDate string
	EndDate   string
}

type ProductOrderHistoryRequest struct {
	AccountID int
	Status    string
}

type OrderProduct struct {
	ProductID       int             `json:"product_id"`
	ProductName     string          `json:"product_name"`
	Quantity        int             `json:"quantity"`
	IndividualPrice decimal.Decimal `json:"individual_price"`
}
type OrdersResponse struct {
	OrderID      int             `json:"order_id"`
	Products     []OrderProduct  `json:"products"`
	TotalPayment decimal.Decimal `json:"total_payment"`
	IsReviewed   bool            `json:"is_reviewed"`
}
