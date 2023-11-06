package repository

import (
	"context"
	"fmt"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	First(ctx context.Context, req dtorepository.ProductRequest) (dtorepository.ProductResponse, error)
	FindProductVariant(ctx context.Context, req dtorepository.FindProductVariantRequest) (dtorepository.FindProductVariantResponse, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) First(ctx context.Context, req dtorepository.ProductRequest) (dtorepository.ProductResponse, error) {
	res := dtorepository.ProductResponse{}

	p := model.Products{}

	qw, err := r.firstWhereQuery(ctx, req)
	if err != nil {
		return res, err
	}

	err = r.db.WithContext(ctx).Where(qw).First(&p).Error
	if err != nil {
		return res, err
	}
	fmt.Println(p)
	res.ID = p.ID
	res.Name = p.Name
	res.Description = p.Description

	return res, nil
}

func (r *productRepository) firstWhereQuery(ctx context.Context, req dtorepository.ProductRequest) (string, error) {
	q := ``

	if req.ProductID != 0 {
		q += fmt.Sprint(` id = ` + fmt.Sprint(req.ProductID))
	}

	return q, nil
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
