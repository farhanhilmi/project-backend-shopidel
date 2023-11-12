package repository

import (
	"context"
	"errors"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	First(ctx context.Context, req dtorepository.ProductRequest) (dtorepository.ProductResponse, error)
	FindProducts(ctx context.Context, req dtorepository.ProductListParam) ([]dtorepository.ProductListResponse, int64, error)
	FindProductVariant(ctx context.Context, req dtorepository.FindProductVariantRequest) (dtorepository.FindProductVariantResponse, error)
	FindProductVariantByID(ctx context.Context, req dtorepository.ProductCart) (dtorepository.ProductCart, error)
	FindProductFavorites(ctx context.Context, req dtorepository.FavoriteProduct) (dtorepository.FavoriteProduct, error)
	FindAllProductFavorites(ctx context.Context, req dtorepository.ProductFavoritesParams) ([]model.FavoriteProductList, int64, error)
	AddProductFavorite(ctx context.Context, req dtorepository.FavoriteProduct) (dtorepository.FavoriteProduct, error)
	RemoveProductFavorite(ctx context.Context, req dtorepository.FavoriteProduct) (dtorepository.FavoriteProduct, error)
	FindSellerAnotherProducts(ctx context.Context, sellerId int) ([]dtousecase.AnotherProduct, error)
	FindProductReviews(ctx context.Context, productId int) ([]dtousecase.ProductReview, error)
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
		distinct on(p.id) p.id,
		p.name,
		aa.district,
		sum(pod.quantity) as total_sold, 
		pvsc.price, 
		pvsc.picture_url, 
		p.created_at,
		p.category_id,
		p.updated_at,
		p.deleted_at
			from products p  
			left join product_variant_selection_combinations pvsc 
				on pvsc.product_id = p.id
			left join account_addresses aa 
				on aa.account_id = p.seller_id  
			left join product_order_details pod 	
				on pod.product_variant_selection_combination_id = pvsc.id
		group by p.id, p.name, pvsc.price, pvsc.picture_url, aa.district
	`

	query := r.db.WithContext(ctx).Table("(?) as t", gorm.Expr(q))
	if req.StartDate != "" {
		query = query.Where("created_at >= ?", req.StartDate)
	}

	if req.EndDate != "" {
		req.EndDate += " 23:59:59"
		query = query.Where("created_at <= ?", req.EndDate)
	}

	if req.CategoryId != "" {
		query = query.Where("category_id = ?", req.CategoryId)
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
	select distinct on (fp.product_id) fp.product_id, sum(pod.quantity) as total_sold, fp.*, p.name, pvsc.price, pvsc.picture_url, aa.district 
		from favorite_products fp 
		left join products p 
			on p.id = fp.product_id
		left join product_variant_selection_combinations pvsc 
			on pvsc.product_id = fp.product_id
		left join account_addresses aa 
			on aa.account_id = fp.account_id  
		left join product_order_details pod 
			on pod.product_variant_selection_combination_id = pvsc.id
		where fp.account_id = ?
	group by fp.product_id, fp.id, p.name, pvsc.price, pvsc.picture_url, aa.district
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
			product_price.lowest_price as "ProductPrice"
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
		where p.seller_id = ?
		limit 12
	`

	err := r.db.Raw(q, sellerId).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *productRepository) FindProductReviews(ctx context.Context, productId int) ([]dtousecase.ProductReview, error) {
	res := []dtousecase.ProductReview{}

	q := `
		select 
			a.full_name as "CustomerName",
			a.profile_picture as "CustomerPictureUrl",
			por.rating as "Stars",
			por.feedback as "Comment",
			'normal' as "Variant",
			por.created_at as "CreatedAt"
		from product_order_reviews por 
		inner join accounts a 
			on a.id = por.account_id 
		where por.product_id = ?
	`

	err := r.db.WithContext(ctx).Raw(q, productId).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
