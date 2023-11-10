package repository

import (
	"context"
	"errors"

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
	FindAllOrderHistoriesByUser(ctx context.Context, req dtorepository.ProductOrderHistoryRequest) ([]model.ProductOrderHistories, error)
	FindAllOrderHistoriesByUserAndStatus(ctx context.Context, req dtorepository.ProductOrderHistoryRequest) ([]model.ProductOrderHistories, error)
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

func (r *productOrdersRepository) FindAllOrderHistoriesByUser(ctx context.Context, req dtorepository.ProductOrderHistoryRequest) ([]model.ProductOrderHistories, error) {
	res := []model.ProductOrderHistories{}

	q := `
	select po.*, pod.quantity, pod.individual_price, pvsc.picture_url, p.name as product_name, pvsc.product_id, 
	por.feedback, por.rating, por.created_at as review_created_at, por.id as review_id
		from product_orders po
		left join product_order_details pod 
			on pod.product_order_id = po.id
		left join product_variant_selection_combinations pvsc 
			on pvsc.id = pod.product_variant_selection_combination_id
		left join products p 
			on p.id = pvsc.product_id 
		left join product_order_reviews por 
			on por.product_order_id = po.id and por.product_id = pvsc.product_id
	where po.account_id = ?
	`
	query := r.db.WithContext(ctx).Table("(?) as t", gorm.Expr(q, req.AccountID))
	if err := query.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (r *productOrdersRepository) FindAllOrderHistoriesByUserAndStatus(ctx context.Context, req dtorepository.ProductOrderHistoryRequest) ([]model.ProductOrderHistories, error) {
	res := []model.ProductOrderHistories{}

	q := `
	select po.*, pod.quantity, pod.individual_price, pvsc.picture_url, p.name as product_name, pvsc.product_id  from product_orders po
		left join product_order_details pod 
			on pod.product_order_id = po.id
		left join product_variant_selection_combinations pvsc 
			on pvsc.id = pod.product_variant_selection_combination_id
		left join products p 
			on p.id = pvsc.product_id 
	where po.account_id = ? and status ilike ?
	`
	query := r.db.WithContext(ctx).Table("(?) as t", gorm.Expr(q, req.AccountID, req.Status))
	if err := query.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
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

	buyerAccount, err := r.accountRepository.RefundBalance(ctx, tx, dtorepository.MyWalletRequest{
		UserID:         req.AccountID,
		Balance:        req.TotalAmount,
		WalletNumber:   req.SellerWalletNumber,
		ProductOrderID: req.ID,
	})
	if err != nil {
		tx.Rollback()
		return res, err
	}

	_, err = r.productVariantCombinationRepository.UpdateStockWithTx(ctx, tx, req.Products)
	if err != nil {
		tx.Rollback()
		return res, err
	}

	_, err = r.accountRepository.DecreaseBalanceSellerWithTx(ctx, tx, dtorepository.MyWalletRequest{UserID: req.SellerID, Balance: req.TotalAmount, WalletNumber: buyerAccount.WalletNumber, ProductOrderID: req.ID})
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

	err := tx.WithContext(ctx).Create(&a).Scan(&res).Error
	if err != nil {
		tx.Rollback()
		return res, err
	}

	orderDetailReq := []model.ProductOrderDetails{}
	productVariants := []model.ProductCombinationVariant{}

	for _, o := range req.ProductVariants {
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

	_, err = r.productDetailRepository.CreateWithTx(ctx, tx, orderDetailReq)
	if err != nil {
		tx.Rollback()
		return res, err
	}

	_, err = r.accountRepository.DecreaseBalanceBuyerWithTx(ctx, tx, dtorepository.MyWalletRequest{
		UserID:          req.AccountID,
		Balance:         req.TotalAmount,
		TransactionType: "Checkout",
		ProductOrderID:  res.ID,
		WalletNumber:    req.SellerWalletNumber,
	})
	if err != nil {
		tx.Rollback()
		return res, err
	}
	_, err = r.accountRepository.IncreaseBalanceSallerWithTx(ctx, tx, dtorepository.MyWalletRequest{
		UserID:          req.SellerID,
		Balance:         req.TotalSellerAmount,
		TransactionType: "Sale",
		ProductOrderID:  res.ID,
		WalletNumber:    req.BuyerWalletNumber,
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
