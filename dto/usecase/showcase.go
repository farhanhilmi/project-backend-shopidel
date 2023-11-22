package dtousecase

import (
	"time"
)

type CreateShowcaseRequest struct {
	ShopId             int
	Name               string
	SelectedProductsId []int
}

type CreateShowcaseResponse struct {
	Id                 int       `json:"id"`
	ShopId             int       `json:"shop_id"`
	Name               string    `json:"name"`
	SelectedProductsId []int     `json:"selected_products_id"`
	CreatedAt          time.Time `json:"created_at"`
}

type GetShowcasesRequest struct {
	ShopId int
	Page   int
}

type GetShowcasesResponse struct {
	Showcases   []Showcase
	CurrentPage int
	TotalPages  int
	TotalItems  int
	Limit       int
}

type Showcase struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetShowcaseDetailResponse struct {
	ID               int                       `json:"id"`
	Name             string                    `json:"name"`
	SelectedProducts []ShowcaseSelectedProduct `json:"selected_products"`
	CreatedAt        time.Time                 `json:"created_at"`
}

type ShowcaseSelectedProduct struct {
	ProductId   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateShowcaseRequest struct {
	Id                 int
	ShopId             int
	Name               string
	SelectedProductsId []int
}

type UpdateShowcaseResponse struct {
	Id                 int       `json:"id"`
	ShopId             int       `json:"shop_id"`
	Name               string    `json:"name"`
	SelectedProductsId []int     `json:"selected_products_id"`
	CreatedAt          time.Time `json:"created_at"`
}
