package dtohttp

type CreateShowcaseRequest struct {
	Name               string `json:"name" binding:"required"`
	SelectedProductsId []int  `json:"selected_products_id" binding:"required"`
}

type UpdateShowcaseRequest struct {
	Name               string `json:"name" binding:"required"`
	SelectedProductsId []int  `json:"selected_products_id" binding:"required"`
}
