package dtousecase

import (
	"mime/multipart"
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
	AccountId int
}

type GetProductDetailRequestV2 struct {
	ShopName    string
	ProductName string
	AccountId   int
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
	ProductReviews []ProductReview  `json:"product_review"`
}

type GetProductSellerResponse struct {
	Id                int                    `json:"id"`
	ProductName       string                 `json:"product_name" binding:"required"`
	Description       string                 `json:"description"`
	HazardousMaterial *bool                  `json:"hazardous_material" binding:"required"`
	IsNew             *bool                  `json:"is_new" binding:"required"`
	InternalSKU       string                 `json:"internal_sku"`
	Weight            decimal.Decimal        `json:"weight" binding:"required"`
	Size              decimal.Decimal        `json:"size" binding:"required"`
	IsActive          *bool                  `json:"is_active" binding:"required"`
	VideoURL          string                 `json:"video_url"`
	CategoryID        int                    `json:"category_id"`
	Images            []string               `json:"images"`
	VariantOptions    []SellerVariantOption  `json:"variant_options"`
	Variants          []ProductVariantSeller `json:"variants"`
}

type ProductCategory struct {
	Id       int                      `json:"id"`
	Name     string                   `json:"name"`
	Children ProductCategoryChildren2 `json:"children,omitempty"`
}

type ProductCategoryChildren2 struct {
	Id       int                      `json:"id"`
	Name     string                   `json:"name"`
	Children ProductCategoryChildren3 `json:"children,omitempty"`
}

type ProductCategoryChildren3 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProductReview struct {
	Id                 int             `json:"id"`
	CustomerName       string          `json:"customer_name"`
	CustomerPictureUrl string          `json:"customer_picture_url"`
	Stars              decimal.Decimal `json:"stars"`
	Comment            string          `json:"comment"`
	Variant            string          `json:"variant,omitempty"`
	CreatedAt          string          `json:"created_at"`
	Pictures           []string        `json:"pictures"`
}

type VariantOption struct {
	VariantOptionName string   `json:"variant_option_name"`
	Childs            []string `json:"childs"`
	Pictures          []string `json:"pictures,omitempty"`
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

type ProductVariantSeller struct {
	ImageURL string                  `json:"image_url"`
	Variant1 ProductSelectionSeller  `json:"variant1"`
	Variant2 *ProductSelectionSeller `json:"variant2,omitempty"`
	Stock    int                     `json:"stock"`
	Price    decimal.Decimal         `json:"price"`
}

type ProductSelectionSeller struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SellerVariantOption struct {
	Name string   `json:"name"`
	Type []string `json:"type"`
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

type ProductListResponse struct {
	ID         int
	Name       string
	District   string
	TotalSold  int
	Price      decimal.Decimal
	Rating     int
	PictureURL string
	ShopName   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type ProductListParam struct {
	AccountID  int
	CategoryId string
	SellerID   int
	SortBy     string
	Search     string
	Sort       string
	MinRating  int
	MinPrice   int
	MaxPrice   int
	District   string
	Limit      int
	Page       int
	StartDate  string
	EndDate    string
}
type ProductOrderHistoryRequest struct {
	AccountID int
	Status    string
	SortBy    string
	Sort      string
	Limit     int
	Page      int
	StartDate string
	EndDate   string
}

type OrderPromotions struct {
	MarketplaceVoucher decimal.Decimal `json:"marketplace_voucher"`
	ShopVoucher        decimal.Decimal `json:"shop_voucher"`
}

type AddressOrder struct {
	Province    string `json:"province"`
	District    string `json:"district"`
	ZipCode     string `json:"zip_code"`
	SubDistrict string `json:"sub_district"`
	Kelurahan   string `json:"kelurahan"`
	Detail      string `json:"detail"`
}
type OrderDetailResponse struct {
	OrderID      int             `json:"order_id"`
	ShopName     string          `json:"shop_name"`
	Status       string          `json:"status"`
	Products     []OrderProduct  `json:"products"`
	TotalPayment decimal.Decimal `json:"total_payment"`
	Promotions   OrderPromotions `json:"promotions"`
	DeliveryFee  decimal.Decimal `json:"delivery_fee"`
	Shipping     AddressOrder    `json:"shipping"`
}

type OrderDetailRequest struct {
	AccountID int
	OrderID   int
}

type ProductOrderReview struct {
	ReviewID       int       `json:"review_id"`
	ReviewFeedback string    `json:"review_feedback"`
	ReviewRating   int       `json:"review_rating"`
	ReviewImageURL string    `json:"review_image_url"`
	CreatedAt      time.Time `json:"created_at"`
}

type OrderProduct struct {
	ProductID            int                `json:"product_id"`
	ProductOrderDetailID int                `json:"product_order_detail_id"`
	ProductName          string             `json:"product_name"`
	VariantName          string             `json:"variant_name"`
	Quantity             int                `json:"quantity"`
	IndividualPrice      decimal.Decimal    `json:"individual_price"`
	Review               ProductOrderReview `json:"review,omitempty"`
	IsReviewed           bool               `json:"is_reviewed"`
}

type SellerOrderProduct struct {
	ProductID            int             `json:"product_id"`
	ProductOrderDetailID int             `json:"product_order_detail_id"`
	ProductName          string          `json:"product_name"`
	VariantName          string          `json:"variant_name"`
	Quantity             int             `json:"quantity"`
	IndividualPrice      decimal.Decimal `json:"individual_price"`
}

type GetProductReviewsRequest struct {
	ProductId int
	Page      int
	Stars     int
	Comment   bool
	Image     bool
	OrderBy   string
	Limit     int
}

type GetProductReviewsResponse struct {
	Reviews     []ProductReview
	TotalPage   int
	TotalItem   int
	CurrentPage int
	Limit       int
}

type ReviewImage struct {
	Url string `json:"url"`
}

type GetProductPicturesRequest struct {
	ProductId int
}

type GetProductPicturesResponse struct {
	PicturesUrl []string
}

type GetProductDetailRecomendedProductRequest struct {
	ProductId int
}

type GetProductDetailRecomendedProductResponse struct {
	AnotherProducts []AnotherProduct `json:"another_products"`
}

type AnotherProduct struct {
	ProductId         int             `json:"product_id"`
	ProductName       string          `json:"product_name"`
	ProductPictureUrl string          `json:"product_picture_url"`
	ProductPrice      decimal.Decimal `json:"product_price"`
	SellerName        string          `json:"seller_name"`
	ProductNameSlug   string          `json:"product_name_slug"`
	ShopNameSlug      string          `json:"shop_name_slug"`
}

type AddNewProductVariantType struct {
	Name  string `form:"name"`
	Value string `form:"value"`
}

type AddNewProductVariant struct {
	Variant1 AddNewProductVariantType `form:"variant1"`
	Variant2 AddNewProductVariantType `form:"variant2"`
	Stock    int                      `form:"stock" binding:"required"`
	Price    decimal.Decimal          `form:"price" binding:"required"`
	ImageID  string                   `form:"image_id" binding:"required"`
	ImageURL string                   `form:"image_url"`
}

type UpdateProductVariant struct {
	Variant1 AddNewProductVariantType `form:"variant1"`
	Variant2 AddNewProductVariantType `form:"variant2"`
	Stock    int                      `form:"stock" binding:"required"`
	Price    decimal.Decimal          `form:"price" binding:"required"`
	ImageID  string                   `form:"image_id"`
	ImageURL string                   `form:"image_url"`
}

type AddNewProductRequest struct {
	ProductID         int
	SellerID          int
	ProductName       string
	Description       string
	DeletedImages     []string
	CategoryID        int
	HazardousMaterial *bool
	IsNew             *bool
	InternalSKU       string
	Weight            decimal.Decimal
	Size              decimal.Decimal
	IsActive          *bool
	Variants          []AddNewProductVariant
	Images            []*multipart.FileHeader
	VideoURL          string
}

type UpdateProductRequest struct {
	ProductID         int
	SellerID          int
	ProductName       string
	Description       string
	DeletedImages     []string
	CategoryID        int
	HazardousMaterial *bool
	IsNew             *bool
	InternalSKU       string
	Weight            decimal.Decimal
	Size              decimal.Decimal
	IsActive          *bool
	Variants          []UpdateProductVariant
	Images            []*multipart.FileHeader
	VideoURL          string
}

type AddNewProductResponse struct {
	ProductName string
}

type UpdatePhoto struct {
	ImageURL    string                `json:"image_url"`
	UserID      int                   `json:"-"`
	Image       multipart.File        `json:"-"`
	ImageHeader *multipart.FileHeader `json:"-"`
}

type UploadNewPhoto struct {
	ImageID     string                `form:"image_id" binding:"required" json:"image_id"`
	Image       multipart.File        `json:"-"`
	ImageHeader *multipart.FileHeader `json:"-"`
}

type ProductVariants struct {
	ID        int
	ProductID int
	Name      string
}

type GetFavoriteProductListResponse struct {
	ID              int             `json:"id"`
	ProductID       int             `json:"product_id"`
	Name            string          `json:"name"`
	District        string          `json:"district"`
	TotalSold       int             `json:"total_sold"`
	Price           decimal.Decimal `json:"price"`
	PictureURL      string          `json:"picture_url"`
	CreatedAt       time.Time       `json:"created_at"`
	ProductNameSlug string          `json:"product_name_slug"`
	ShopNameSlug    string          `json:"shop_name_slug"`
}
