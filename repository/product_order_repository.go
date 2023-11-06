package repository

import (
	"context"
	"errors"
	"log"

	dtorepository "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productOrdersRepository struct {
	db                                  *gorm.DB
	accountRepository                   accountRepository
	productVariantCombinationRepository productVariantCombinationRepository
	productDetailRepository             productDetailRepository
}

type ProductOrdersRepository interface {
	FindByIDAndAccountID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error)
	FindByStatusAndAccountID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error)
	UpdateOrderStatusByIDAndAccountID(ctx context.Context, req dtorepository.ReceiveOrderRequest) (dtorepository.ProductOrderResponse, error)
	FindByIDAndSellerID(ctx context.Context, req dtorepository.ProductOrderRequest) ([]model.ProductOrderSeller, error)
	ProcessedOrder(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error)
	Create(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error)
}

func NewProductOrdersRepository(db *gorm.DB) ProductOrdersRepository {
	return &productOrdersRepository{
		db: db,
	}
}

func (r *productOrdersRepository) FindByID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error) {
	order := model.ProductOrders{}
	res := dtorepository.ProductOrderResponse{}

	err := r.db.WithContext(ctx).Where("id = ?", req.ID).First(&order).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	res.ID = order.ID
	res.AccountID = order.AccountID
	res.CourierID = order.CourierID
	res.DeliveryFee = order.DeliveryFee
	res.Province = order.Province
	res.SubDistrict = order.SubDistrict
	res.Kelurahan = order.Kelurahan
	res.Status = order.Status
	res.ZipCode = order.ZipCode
	res.District = order.District
	res.CreatedAt = order.CreatedAt
	res.DeletedAt = order.DeletedAt
	res.UpdatedAt = order.UpdatedAt

	return res, err
}

func (r *productOrdersRepository) FindByIDAndSellerID(ctx context.Context, req dtorepository.ProductOrderRequest) ([]model.ProductOrderSeller, error) {
	order := []model.ProductOrderSeller{}

	err := r.db.WithContext(ctx).Raw(`
	select po.id, po.account_id, p.id as product_id, p.seller_id, pod.individual_price, pod.quantity, po.status, pvsc.id as product_variant_selection_combination_id, pvsc.stock as product_stock
		from product_order_details pod
		left join product_variant_selection_combinations pvsc on pod.product_variant_selection_combination_id = pvsc.id 
		left join products p on p.id = pvsc.product_id 
		left join product_orders po on po.id = pod.product_order_id 
	where pod.product_order_id = ? and p.seller_id = ? and po.status = ?;
	`, req.ID, req.SellerID, req.Status).Scan(&order).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return order, util.ErrNoRecordFound
	}

	return order, err
}

func (r *productOrdersRepository) FindByIDAndAccountID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error) {
	order := model.ProductOrders{}
	res := dtorepository.ProductOrderResponse{}

	err := r.db.WithContext(ctx).Where("account_id = ?", req.AccountID).Where("id = ?", req.ID).First(&order).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	res.ID = order.ID
	res.AccountID = order.AccountID
	res.CourierID = order.CourierID
	res.DeliveryFee = order.DeliveryFee
	res.Province = order.Province
	res.SubDistrict = order.SubDistrict
	res.Kelurahan = order.Kelurahan
	res.Status = order.Status
	res.ZipCode = order.ZipCode
	res.District = order.District
	res.CreatedAt = order.CreatedAt
	res.DeletedAt = order.DeletedAt
	res.UpdatedAt = order.UpdatedAt

	return res, err
}

func (r *productOrdersRepository) FindByStatusAndAccountID(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error) {
	order := model.ProductOrders{}
	res := dtorepository.ProductOrderResponse{}

	err := r.db.WithContext(ctx).Where("account_id = ?", req.AccountID).Where("status = ?", req.Status).First(&order).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, util.ErrNoRecordFound
	}

	res.ID = order.ID
	res.AccountID = order.AccountID
	res.CourierID = order.CourierID
	res.DeliveryFee = order.DeliveryFee
	res.Province = order.Province
	res.SubDistrict = order.SubDistrict
	res.Kelurahan = order.Kelurahan
	res.Status = order.Status
	res.ZipCode = order.ZipCode
	res.District = order.District
	res.CreatedAt = order.CreatedAt
	res.DeletedAt = order.DeletedAt
	res.UpdatedAt = order.UpdatedAt

	return res, err
}

func (r *productOrdersRepository) UpdateOrderStatusByIDAndAccountID(ctx context.Context, req dtorepository.ReceiveOrderRequest) (dtorepository.ProductOrderResponse, error) {
	order := model.ProductOrders{}
	res := dtorepository.ProductOrderResponse{}
	tx := r.db.Begin()

	err := tx.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(&order).Where("id = ?", req.ID).Update("status", req.Status).Error

	if err != nil {
		tx.Rollback()
		return res, err
	}

	_, err = r.accountRepository.RefundBalance(ctx, tx, dtorepository.MyWalletRequest{UserID: req.AccountID, Balance: req.TotalAmount})
	if err != nil {
		tx.Rollback()
		return res, err
	}

	_, err = r.productVariantCombinationRepository.UpdateStockWithTx(ctx, tx, req.Products)
	if err != nil {
		tx.Rollback()
		return res, err
	}

	_, err = r.accountRepository.DecreaseBalanceSellerWithTx(ctx, tx, dtorepository.MyWalletRequest{UserID: req.SellerID, Balance: req.TotalAmount})
	if err != nil {
		tx.Rollback()
		return res, err
	}
	tx.Commit()

	res.ID = order.ID
	res.AccountID = order.AccountID
	res.CourierID = order.CourierID
	res.DeliveryFee = order.DeliveryFee
	res.Province = order.Province
	res.SubDistrict = order.SubDistrict
	res.Kelurahan = order.Kelurahan
	res.Status = order.Status
	res.ZipCode = order.ZipCode
	res.District = order.District
	res.CreatedAt = order.CreatedAt
	res.DeletedAt = order.DeletedAt
	res.UpdatedAt = order.UpdatedAt

	return res, nil
}

func (r *productOrdersRepository) ProcessedOrder(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error) {
	order := model.ProductOrders{}
	res := dtorepository.ProductOrderResponse{}

	err := r.db.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Table: clause.Table{
			Name: clause.CurrentTable,
		}}).Model(&order).Where("id = ?", req.ID).Update("status", req.Status).Error

	if err != nil {
		return res, err
	}

	res.ID = order.ID
	res.AccountID = order.AccountID
	res.CourierID = order.CourierID
	res.DeliveryFee = order.DeliveryFee
	res.Province = order.Province
	res.SubDistrict = order.SubDistrict
	res.Kelurahan = order.Kelurahan
	res.Status = order.Status
	res.ZipCode = order.ZipCode
	res.District = order.District
	res.CreatedAt = order.CreatedAt
	res.DeletedAt = order.DeletedAt
	res.UpdatedAt = order.UpdatedAt

	return res, nil
}

func (r *productOrdersRepository) Create(ctx context.Context, req dtorepository.ProductOrderRequest) (dtorepository.ProductOrderResponse, error) {
	res := dtorepository.ProductOrderResponse{}
	tx := r.db.Begin()

	a := model.ProductOrders{
		CourierID:     req.CourierID,
		AccountID:     req.AccountID,
		DeliveryFee:   req.DeliveryFee,
		District:      req.District,
		Province:      req.Province,
		SubDistrict:   req.SubDistrict,
		Kelurahan:     req.Kelurahan,
		Notes:         req.Notes,
		Status:        req.Status,
		ZipCode:       req.ZipCode,
		AddressDetail: req.AddressDetail,
	}

	err := tx.WithContext(ctx).Create(&a).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	orderDetailReq := []model.ProductOrderDetails{}
	productVariants := []model.ProductCombinationVariant{}

	for _, o := range req.ProductVariants {
		log.Println("O", o)
		variant := model.ProductCombinationVariant{
			ID:    o.ProductVariantSelectionCombinationID,
			Stock: o.Quantity,
		}
		product := model.ProductOrderDetails{
			ProductOrderID:                       res.ID,
			ProductVariantSelectionCombinationID: o.ProductVariantSelectionCombinationID,
			Quantity:                             o.Quantity,
			IndividualPrice:                      o.IndividualPrice,
		}
		productVariants = append(productVariants, variant)
		orderDetailReq = append(orderDetailReq, product)
	}
	log.Println("orderDetailReq", orderDetailReq)

	_, err = r.productDetailRepository.CreateWithTx(ctx, tx, orderDetailReq)
	if err != nil {
		tx.Rollback()
		return res, err
	}

	_, err = r.accountRepository.DecreaseBalanceBuyerWithTx(ctx, tx, dtorepository.MyWalletRequest{
		UserID:          req.AccountID,
		Balance:         req.TotalAmount,
		TransactionType: "Checkout",
	})
	if err != nil {
		tx.Rollback()
		return res, err
	}
	_, err = r.accountRepository.IncreaseBalanceSallerWithTx(ctx, tx, dtorepository.MyWalletRequest{
		UserID:          req.AccountID,
		Balance:         req.TotalAmount,
		TransactionType: "Sale",
	})
	if err != nil {
		tx.Rollback()
		return res, err
	}

	_, err = r.productVariantCombinationRepository.DecreaseStockWithTx(ctx, tx, productVariants)
	if err != nil {
		tx.Rollback()
		return res, err
	}

	tx.Commit()

	res.ID = a.ID
	res.AccountID = a.AccountID
	res.CourierID = a.CourierID
	res.DeliveryFee = a.DeliveryFee
	res.Province = a.Province
	res.SubDistrict = a.SubDistrict
	res.Kelurahan = a.Kelurahan
	res.Status = a.Status
	res.ZipCode = a.ZipCode
	res.District = a.District
	res.CreatedAt = a.CreatedAt
	res.DeletedAt = a.DeletedAt
	res.UpdatedAt = a.UpdatedAt

	return res, err
}
