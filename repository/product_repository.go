package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	First(ctx context.Context, req dtorepository.ProductRequest) (dtorepository.ProductResponse, error)
	FirstV2(ctx context.Context, req dtorepository.ProductRequestV2) (dtorepository.ProductResponse, error)
	FindProducts(ctx context.Context, req dtorepository.ProductListParam) ([]dtorepository.ProductListResponse, int64, error)
	FindImages(ctx context.Context, productId int) (dtorepository.FindProductPicturesResponse, error)
	FindProductVariant(ctx context.Context, req dtorepository.FindProductVariantRequest) (dtorepository.FindProductVariantResponse, error)
	FindProductVariantByID(ctx context.Context, req dtorepository.ProductCart) (dtorepository.ProductCart, error)
	FindProductFavorites(ctx context.Context, req dtorepository.FavoriteProduct) (dtorepository.FavoriteProduct, error)
	FindAllProductFavorites(ctx context.Context, req dtorepository.ProductFavoritesParams) ([]model.FavoriteProductList, int64, error)
	AddProductFavorite(ctx context.Context, req dtorepository.FavoriteProduct) (dtorepository.FavoriteProduct, error)
	RemoveProductFavorite(ctx context.Context, req dtorepository.FavoriteProduct) (dtorepository.FavoriteProduct, error)
	FindSellerAnotherProducts(ctx context.Context, sellerId int) ([]dtousecase.AnotherProduct, error)
	FindProductReviews(ctx context.Context, req dtousecase.GetProductReviewsRequest) (dtorepository.GetProductReviewsResponse, error)
	FindByID(ctx context.Context, req dtorepository.ProductRequest) (dtorepository.ProductResponse, error)
	AddNewProduct(ctx context.Context, req dtorepository.AddNewProductRequest) (dtorepository.AddNewProductResponse, error)
	FindProductReviewPictures(ctx context.Context, reviewId int) ([]dtousecase.ReviewImage, error)
	RemoveProductByID(ctx context.Context, req dtorepository.RemoveProduct) (dtorepository.RemoveProduct, error)
	FindByIDAndSeller(ctx context.Context, req dtorepository.ProductRequest) (dtorepository.ProductResponse, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) FindProducts(ctx context.Context, req dtorepository.ProductListParam) ([]dtorepository.ProductListResponse, int64, error) {
	res := []dtorepository.ProductListResponse{}
	var totalItems int64

	q := `
		select
			p.id,
			p.name,
			p.description,
			aa.district,
			0 as total_sold, 
			product_price.lowest_price as "price", 
			round( CAST(float8 (random() * 5) as numeric), 1) as rating,
			product_image.picture_url, 
			case
				when 0 in ($1) then
					case 
						when category_level_2.level_2_id is not null then category_level_2.level_2_id
						when category_level_3.level_2_id is not null then category_level_3.level_2_id
					end
				else 
					case 
						when category_level_1.level_1_id is not null then category_level_1.level_1_id
						when category_level_2.level_2_id is not null then category_level_2.level_2_id
						when category_level_3.level_3_id is not null then category_level_3.level_3_id
					end
			end as "category_id",
			case
				when 0 in ($1) then
					case
						when category_level_2.level_2_name is not null then category_level_2.level_2_name
						when category_level_3.level_2_name is not null then category_level_3.level_2_name
					end
				else
					case
						when category_level_1.level_1_name is not null then category_level_1.level_1_name
						when category_level_2.level_2_name is not null then category_level_2.level_2_name
						when category_level_3.level_3_name is not null then category_level_3.level_3_name
					end
			end as "category_name",
			p.created_at,
			p.updated_at,
			p.deleted_at,
			seller.shop_name,
			TRIM(BOTH '-' FROM 
				REGEXP_REPLACE(
					REGEXP_REPLACE(
						LOWER(p."name"), 
						'[^a-z0-9]+', '-', 'g'
					), 
					'-+', '-', 'g'
				)
			) AS "ProductNameSlug",
			TRIM(BOTH '-' FROM 
				REGEXP_REPLACE(
					REGEXP_REPLACE(
						LOWER(seller.shop_name), 
						'[^a-z0-9]+', '-', 'g'
					), 
					'-+', '-', 'g'
				)
			) AS "ShopNameSlug"
		from products p
			inner join lateral (
					select	
						pi2.product_id,
						pi2.url as picture_url
					from product_images pi2 
					where pi2.product_id = p.id 
					order by pi2.id asc
					limit 1
				) product_image on product_image.product_id = p.id 
			inner join lateral (
					select
						pvsc2.product_id,
						min(pvsc2.price) as lowest_price
					from product_variant_selection_combinations pvsc2
					where pvsc2.product_id = p.id
					group by pvsc2.product_id 
				) product_price on product_price.product_id = p.id
			left join account_addresses aa 
				on aa.account_id = p.seller_id
				and aa.is_seller_default is true
			left join accounts as seller
				on seller.id = p.seller_id
			left join (
				select
					c.id as level_1_id,
					c."name" as level_1_name
				from categories c
				where c."level" = 1
			) as category_level_1 on category_level_1.level_1_id = p.category_id 
			left join (
				select
					c.id as level_2_id,
					c."name" level_2_name,
					c2.id as level_1_id,
					c2."name" as level_1_name
				from categories c
				inner join categories c2 
					on c2.id = c.parent 
				where c."level" = 2
			) as category_level_2 on category_level_2.level_2_id = p.category_id 
			left join (
				select
					c.id as level_3_id,
					c."name" level_3_name,
					c2.id as level_2_id,
					c2."name" level_2_name,
					c3.id as level_1_id,
					c3."name" as level_1_name
				from categories c
				inner join categories c2 
					on c2.id = c.parent 
				inner join categories c3
					on c3.id = c2.parent 
				where c."level" = 3
			) as category_level_3 on category_level_3.level_3_id = p.category_id
			left join (
				select 
					distinct(a.id) as id
				from (
					(
						select
							level_3.id 
						from categories level_3
						left join categories level_2
							on level_2.id = level_3.parent 
							and level_2."level" = 2
						left join categories level_1
							on level_1.id = level_2.parent 
							and level_1."level" = 1
						where level_3."level" = 3
							and (
								level_3.id in ($1)
								or level_2.id in ($1)
								or level_1.id in ($1)
							)
					) union all (
						select
							level_2.id 
						from categories level_2
						left join categories level_1
							on level_1.id = level_2.parent
							and level_1."level" = 1
						where level_2."level" = 2
							and (
								level_2.id in ($1)
								or level_1.id in ($1)
							) 
					) union all (
						select
							level_1.id 
						from categories level_1
						where level_1."level" = 1
							and level_1.id in ($1)
					)
				) a
			) as child on child.id = p.category_id 
			where 
				(
					0 in ($1)
					or child.id is not null
				)
	`

	categoriesId := []string{"0"}
	if req.CategoryId != "" && !strings.Contains(req.CategoryId, "#") {
		categoriesId = []string{req.CategoryId}
	} else if req.CategoryId != "" && strings.Contains(req.CategoryId, "#") {
		categoriesId = strings.Split(req.CategoryId, "#")
	}

	query := r.db.WithContext(ctx).Table("(?) as t", r.db.Raw(q, categoriesId))
	if req.StartDate != "" {
		query = query.Where("created_at >= ?", req.StartDate)
	}

	if req.EndDate != "" {
		req.EndDate += " 23:59:59"
		query = query.Where("created_at <= ?", req.EndDate)
	}

	if req.MinRating > 0 && req.MinRating <= 5 {
		query = query.Where("rating >= ?", req.MinRating)
	}

	if req.MinPrice > 0 {
		query = query.Where("price >= ?", req.MinPrice)
	}

	if req.MaxPrice > 0 {
		query = query.Where("price <= ?", req.MaxPrice)
	}

	if req.District != "" && !strings.Contains(req.District, "#") {
		query = query.Where("district ilike ?", req.District)
	} else if req.District != "" && strings.Contains(req.District, "#") {
		districts := strings.Split(req.District, "#")
		query = query.Where("district IN ?", districts)
	}

	if req.Search != "" {
		find := "%" + req.Search + "%"
		query = query.
			Where(
				"name ilike ? or description ilike ?",
				find,
				find,
			)
	}

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	query = query.Order(req.SortBy + " " + req.Sort)
	offset := (req.Page - 1) * req.Limit
	query = query.Offset(offset).Limit(req.Limit)

	if err := query.Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, totalItems, nil
}

func (r *productRepository) FindProductVariantByID(ctx context.Context, req dtorepository.ProductCart) (dtorepository.ProductCart, error) {
	account := model.ProductVariantSelectionCombinations{}

	err := r.db.WithContext(ctx).Model(&account).Where("id = ?", req.ProductID).First(&account).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dtorepository.ProductCart{}, util.ErrNoRecordFound
	}
	if err != nil {
		return dtorepository.ProductCart{}, err
	}

	return dtorepository.ProductCart{
		ProductID: req.ProductID,
		Quantity:  account.Stock,
		ID:        account.ID,
	}, nil
}

func (r *productRepository) FindByID(ctx context.Context, req dtorepository.ProductRequest) (dtorepository.ProductResponse, error) {
	res := dtorepository.ProductResponse{}

	err := r.db.WithContext(ctx).Model(&model.Products{}).Where("id = ?", req.ProductID).First(&model.Products{}).Scan(&res).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *productRepository) FindByIDAndSeller(ctx context.Context, req dtorepository.ProductRequest) (dtorepository.ProductResponse, error) {
	res := dtorepository.ProductResponse{}

	err := r.db.WithContext(ctx).Model(&model.Products{}).Where("id = ?", req.ProductID).Where("seller_id", req.AccountId).First(&model.Products{}).Scan(&res).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *productRepository) First(ctx context.Context, req dtorepository.ProductRequest) (dtorepository.ProductResponse, error) {
	res := dtorepository.ProductResponse{}

	q := `
		select 
			p.Id as "ID",
			p."name" as "Name",
			p.description as "Description",
			case 
				when fp.id is not null then true
				else false
			end as "IsFavorite",
			p.seller_id as "SellerId"
		from products p 
		left join favorite_products fp 
			on fp.product_id = p.id 
			and fp.account_id = $1
		where p.Id = $2
	`

	err := r.db.WithContext(ctx).Raw(q, req.AccountId, req.ProductID).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *productRepository) FirstV2(ctx context.Context, req dtorepository.ProductRequestV2) (dtorepository.ProductResponse, error) {
	res := dtorepository.ProductResponse{}

	q := `
		select 
			p.Id as "ID",
			p."name" as "Name",
			p.description as "Description",
			case 
				when fp.id is not null then true
				else false
			end as "IsFavorite",
			p.seller_id as "SellerId"
		from products p 
		left join favorite_products fp 
			on fp.product_id = p.id 
			and fp.account_id = $1
		inner join accounts a 
			on a.id = p.seller_id 
			and TRIM(BOTH '-' FROM 
					REGEXP_REPLACE(
						REGEXP_REPLACE(
							LOWER(a.shop_name), 
							'[^a-z0-9]+', '-', 'g'
						), 
						'-+', '-', 'g'
					)
				) ilike $2
		where TRIM(BOTH '-' FROM 
				REGEXP_REPLACE(
					REGEXP_REPLACE(
						LOWER(p.name), 
						'[^a-z0-9]+', '-', 'g'
					), 
					'-+', '-', 'g'
				)
			) ilike $3
	`

	err := r.db.WithContext(ctx).Raw(q, req.AccountId, req.ShopName, req.ProductName).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *productRepository) FindImages(ctx context.Context, productId int) (dtorepository.FindProductPicturesResponse, error) {
	res := dtorepository.FindProductPicturesResponse{}
	pictures := []dtorepository.ProductPicture{}

	q := `
		select
			pi2.url as "PictureUrl"
		from product_images pi2 
		where pi2.product_id = $1
	`

	err := r.db.WithContext(ctx).Raw(q, productId).Scan(&pictures).Error
	if err != nil {
		return res, err
	}

	res.ProductPictures = pictures

	return res, nil
}

func (r *productRepository) FindProductVariant(ctx context.Context, req dtorepository.FindProductVariantRequest) (dtorepository.FindProductVariantResponse, error) {
	res := dtorepository.FindProductVariantResponse{}
	variants := []dtorepository.ProductVariant{}

	q := `
		select 
			pvsc.id as "VariantId",
			pvsc.product_variant_selection_id1 as "SelectionId1",
			pvs."name" as "SelectionName1",
			pv."name" as "VariantName1",
			pvsc.product_variant_selection_id2 as "SelectionId2",
			pvs2."name" as "SelectionName2",
			pv2."name" as "VariantName2",
			pvsc.price, 
			pvsc.stock
		from product_variant_selection_combinations pvsc
		left join product_variant_selections pvs
			on pvs.id = pvsc.product_variant_selection_id1
		left join product_variants pv
			on pv.id = pvs.product_variant_id
		left join product_variant_selections pvs2 
			on pvs2.id = pvsc.product_variant_selection_id2 
		left join product_variants pv2
			on pv2.id = pvs2.product_variant_id
		where pvsc.product_id = ?
		order by pvs.id asc, pvs2.id asc
	`

	err := r.db.WithContext(ctx).Raw(q, req.ProductId).Scan(&variants).Error
	if err != nil {
		return res, err
	}

	res.Variants = variants

	return res, nil
}

func (r *productRepository) AddProductFavorite(ctx context.Context, req dtorepository.FavoriteProduct) (dtorepository.FavoriteProduct, error) {
	res := dtorepository.FavoriteProduct{}

	err := r.db.WithContext(ctx).Model(&model.FavoriteProducts{}).Create(&req).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (r *productRepository) RemoveProductFavorite(ctx context.Context, req dtorepository.FavoriteProduct) (dtorepository.FavoriteProduct, error) {
	res := dtorepository.FavoriteProduct{}

	err := r.db.WithContext(ctx).Where("account_id = ?", req.AccountID).Where("product_id = ?", req.ProductID).Delete(&model.FavoriteProducts{}).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (r *productRepository) FindProductFavorites(ctx context.Context, req dtorepository.FavoriteProduct) (dtorepository.FavoriteProduct, error) {
	res := dtorepository.FavoriteProduct{}

	err := r.db.WithContext(ctx).Model(&model.FavoriteProducts{}).Where("product_id = ?", req.ProductID).Where("account_id = ?", req.AccountID).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}
	if err != nil {
		return res, err
	}

	return res, err
}

func (r *productRepository) FindAllProductFavorites(ctx context.Context, req dtorepository.ProductFavoritesParams) ([]model.FavoriteProductList, int64, error) {
	res := []model.FavoriteProductList{}
	var totalItems int64

	q := `
		select 
			distinct on (fp.product_id) fp.product_id, 
			0 as total_sold, 
			fp.*, 
			p.name, 
			product_price.lowest_price as price, 
			product_image.picture_url, 
			aa.district 
		from favorite_products fp 
			left join products p 
					on p.id = fp.product_id
			left join product_variant_selection_combinations pvsc 
					on pvsc.product_id = fp.product_id
			left join account_addresses aa 
					on aa.account_id = fp.account_id 
			inner join lateral (
					select	
						pi2.product_id,
						pi2.url as picture_url
					from product_images pi2 
					where pi2.product_id = p.id 
					order by pi2.id asc
					limit 1
				) product_image on product_image.product_id = p.id 
			inner join lateral (
					select
						pvsc2.product_id,
						min(pvsc2.price) as lowest_price
					from product_variant_selection_combinations pvsc2
					where pvsc2.product_id = p.id
					group by pvsc2.product_id 
				) product_price on product_price.product_id = p.id
		where fp.account_id = ?
	`
	query := r.db.WithContext(ctx).Table("(?) as t", gorm.Expr(q, req.AccountID))

	if req.StartDate != "" {
		query = query.Where("created_at >= ?", req.StartDate)
	}

	if req.EndDate != "" {
		req.EndDate += " 23:59:59"
		query = query.Where("created_at <= ?", req.EndDate)
	}

	if req.Search != "" {
		query = query.Where("name ilike ?", "%"+req.Search+"%")
	}

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	query = query.Order(req.SortBy + " " + req.Sort)
	offset := (req.Page - 1) * req.Limit
	query = query.Offset(offset).Limit(req.Limit)

	if err := query.Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, totalItems, nil
}

func (r *productRepository) FindSellerAnotherProducts(ctx context.Context, sellerId int) ([]dtousecase.AnotherProduct, error) {
	res := []dtousecase.AnotherProduct{}

	q := `
		select
			p.id as "ProductId",
			p.name as "ProductName",
			product_image.url as "ProductPictureUrl",
			product_price.lowest_price as "ProductPrice",
			seller.shop_name as "SellerName",
			TRIM(BOTH '-' FROM 
				REGEXP_REPLACE(
					REGEXP_REPLACE(
						LOWER(p.name), 
						'[^a-z0-9]+', '-', 'g'
					), 
					'-+', '-', 'g'
				)
			) as "ProductNameSlug",
			TRIM(BOTH '-' FROM 
				REGEXP_REPLACE(
					REGEXP_REPLACE(
						LOWER(seller.shop_name), 
						'[^a-z0-9]+', '-', 'g'
					), 
					'-+', '-', 'g'
				)
			) as "ShopNameSlug"
		from products p
		left join lateral (
			select
				pi2.product_id,
				url 
			from product_images pi2
			where pi2.product_id = p.id 
			order by pi2.id asc
			limit 1
		) product_image on product_image.product_id = p.id 
		left join (
			select
				pvsc.product_id,
				min(pvsc.price) as lowest_price
			from product_variant_selection_combinations pvsc 
			group by pvsc.product_id 
		) product_price on product_price.product_id = p.id 
		left join accounts seller
			on seller.id = p.seller_id
		where p.seller_id = ?
		limit 12
	`

	err := r.db.Raw(q, sellerId).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *productRepository) FindProductReviews(ctx context.Context, req dtousecase.GetProductReviewsRequest) (dtorepository.GetProductReviewsResponse, error) {
	res := dtorepository.GetProductReviewsResponse{}
	pr := []dtorepository.ProductReview{}
	pg := dtogeneral.PaginationData{}

	q := `
		select 
			por.id as "Id",
			a.full_name as "CustomerName",
			a.profile_picture as "CustomerPictureUrl",
			por.rating as "Stars",
			por.feedback as "Comment",
			'normal' as "Variant",
			por.created_at as "CreatedAt"
		from product_order_reviews por 
		inner join accounts a 
			on a.id = por.account_id 
	`

	pq := `
		select 
			count(por.id) as "TotalItem"
		from product_order_reviews por 
		inner join accounts a 
			on a.id = por.account_id 
	`

	wq := ` where por.product_id = $1 `

	oq := ` order by por.created_at ` + req.OrderBy

	lq := ` limit ` + fmt.Sprint(req.Limit, " ")

	ofq := ` offset ` + fmt.Sprint((req.Page-1)*req.Limit, " ")

	if req.Stars != 0 {
		wq += fmt.Sprint(` and por.rating::int = `, req.Stars, " ")
	}
	if req.Comment {
		wq += ` and por.feedback is not null `
	}

	err := r.db.WithContext(ctx).Raw(q+wq+oq+lq+ofq, req.ProductId).Scan(&pr).Error
	if err != nil {
		return res, err
	}

	err = r.db.WithContext(ctx).Raw(pq+wq, req.ProductId).Scan(&pg).Error
	if err != nil {
		return res, err
	}

	res.Reviews = pr
	res.TotalItem = pg.TotalItem
	res.Limit = req.Limit
	res.TotalPage = int(math.Ceil(float64(pg.TotalItem) / float64(res.Limit)))
	res.CurrentPage = req.Page

	return res, nil
}

func removeDuplicateValues(arr []dtorepository.ProductVariantType) []dtorepository.ProductVariantType {
	length := len(arr) - 1
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			if arr[i].VariantValue == arr[j].VariantValue {
				arr[j] = arr[length]
				arr = arr[0:length]
				length--
				j--
			}
		}
	}

	return arr
}

func (r *productRepository) AddNewProduct(ctx context.Context, req dtorepository.AddNewProductRequest) (dtorepository.AddNewProductResponse, error) {
	res := dtorepository.AddNewProductResponse{}

	product := model.Products{
		Name:              req.ProductName,
		Description:       req.Description,
		CategoryID:        req.CategoryID,
		SellerID:          req.SellerID,
		HazardousMaterial: *req.HazardousMaterial,
		Weight:            req.Weight,
		Size:              req.Size,
		IsNew:             *req.IsNew,
		IsActive:          *req.IsActive,
		InternalSKU:       req.InternalSKU,
	}

	tx := r.db.Begin()

	err := tx.WithContext(ctx).Model(&model.Products{}).Create(&product).Scan(&res).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	if req.VideoURL != "" {
		err = tx.WithContext(ctx).Create(&model.ProductVideos{ProductID: res.ID, URL: req.VideoURL}).Error
		if err != nil {
			tx.Rollback()
			return res, err
		}
	}

	productImages := []model.ProductImages{}

	for _, url := range req.Images {
		image := model.ProductImages{
			ProductID: res.ID,
			URL:       url,
		}
		productImages = append(productImages, image)
	}

	err = tx.WithContext(ctx).Create(&productImages).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	productVariants := []model.ProductVariants{}

	for _, v := range req.ProductVariants {
		variant := model.ProductVariants{
			ProductID: res.ID,
			Name:      v.Name,
		}
		productVariants = append(productVariants, variant)
	}

	log.Println(productVariants)

	err = tx.WithContext(ctx).Create(&productVariants).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	variantSelections := []dtorepository.ProductVariantType{}

	for _, v := range req.Variants {
		selections := []dtorepository.ProductVariantType{
			{
				VariantName:  v.Variant1.Name,
				VariantValue: v.Variant1.Value,
			},
			{
				VariantName:  v.Variant2.Name,
				VariantValue: v.Variant2.Value,
			},
		}

		variantSelections = append(variantSelections, selections...)
	}

	selections := removeDuplicateValues(variantSelections)

	productVariantSelections := []model.ProductVariantSelections{}

	for _, selection := range selections {
		for _, v := range productVariants {
			if v.Name == selection.VariantName {
				variant := model.ProductVariantSelections{
					ProductVariantID: v.ID,
					Name:             selection.VariantValue,
				}
				productVariantSelections = append(productVariantSelections, variant)
			}
		}
	}

	err = tx.WithContext(ctx).Create(&productVariantSelections).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	for _, v := range req.Variants {
		var imageUrl string
		if v.ImageID != "" && v.Variant1.Name != "" {
			imageUrl, err = util.GetVariantImageURL(v.ImageID)
			if err != nil {
				tx.Rollback()
				return res, err
			}
		}

		variantCombination := model.ProductVariantSelectionCombinations{
			ProductID:  res.ID,
			Price:      v.Price,
			PictureURL: imageUrl,
			Stock:      v.Stock,
		}
		for _, selection := range productVariantSelections {
			if v.Variant1.Value == selection.Name {
				variantCombination.ProductVariantSelectionID1 = selection.ID
			}
		}
		for _, selection := range productVariantSelections {
			if v.Variant2.Value == selection.Name {
				variantCombination.ProductVariantSelectionID2 = selection.ID
			}
		}
		err = tx.WithContext(ctx).Create(&variantCombination).Error
		if err != nil {
			tx.Rollback()
			return res, err
		}
	}

	tx.Commit()

	return res, err
}

func (r *productRepository) FindProductReviewPictures(ctx context.Context, reviewId int) ([]dtousecase.ReviewImage, error) {
	res := []dtousecase.ReviewImage{}

	q := `
		select 
			pori.image_url as url
		from product_order_review_images pori 
		where pori.product_review_id = ?
	`

	err := r.db.WithContext(ctx).Raw(q, reviewId).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *productRepository) RemoveProductByID(ctx context.Context, req dtorepository.RemoveProduct) (dtorepository.RemoveProduct, error) {
	res := dtorepository.RemoveProduct{}

	tx := r.db.Begin()

	err := tx.WithContext(ctx).Model(&model.Products{}).Where("id = ?", req.ID).Where("seller_id = ?", req.SellerID).Update("deleted_at", time.Now()).Scan(&res).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	err = tx.WithContext(ctx).Where("product_id = ?", req.ID).Delete(&model.ProductImages{}).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	err = tx.WithContext(ctx).Where("product_id = ?", req.ID).Delete(&model.ProductVideos{}).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	err = tx.WithContext(ctx).Where("product_id = ?", req.ID).Delete(&model.FavoriteProducts{}).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	showcaseProduct := model.SellerShowcaseProduct{}
	err = tx.WithContext(ctx).Where("product_id = ?", req.ID).Delete(&showcaseProduct).Scan(&showcaseProduct).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	err = tx.WithContext(ctx).Where("id = ?", showcaseProduct.ID).Delete(&model.SellerShowcase{}).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	productVariant := model.ProductVariants{}
	err = tx.WithContext(ctx).Clauses(clause.Returning{}).Where("product_id = ?", req.ID).Delete(&productVariant).Scan(&productVariant).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	productVariantSelection := []model.ProductVariantSelections{}
	err = tx.WithContext(ctx).Clauses(clause.Returning{}).Where("product_variant_id = ?", productVariant.ID).Delete(&productVariantSelection).Scan(&productVariantSelection).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	variantSelectionIds := []int{}
	for _, s := range productVariantSelection {
		variantSelectionIds = append(variantSelectionIds, s.ID)
	}

	productVariantCombinations := []model.ProductVariantSelectionCombinations{}
	err = tx.WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("product_variant_selection_id1 IN ?", variantSelectionIds).
		Or("product_variant_selection_id2 IN ?", variantSelectionIds).
		Delete(&productVariantCombinations).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	variantSelectionCombinationIds := []int{}
	for _, s := range productVariantSelection {
		variantSelectionCombinationIds = append(variantSelectionCombinationIds, s.ID)
	}

	err = tx.WithContext(ctx).Where("product_variant_selection_combination_id IN ?", variantSelectionCombinationIds).
		Delete(&model.AccountCarts{}).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	tx.Commit()

	return res, err
}
