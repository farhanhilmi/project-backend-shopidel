package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/config"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/router"
	router_seller "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/router/seller"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Start(gin *gin.Engine, db *gorm.DB) {
	gin.Use(middleware.ErrorHandler())

	accountRepo := repository.NewAccountRepository(db)
	usedEmailRepo := repository.NewUsedEmailRepository(db)
	productOrderRepo := repository.NewProductOrdersRepository(db)
	productVariantCombinationRepo := repository.NewProductVariantCombinationRepository(db)
	accountAddressRepo := repository.NewAccountAddressRepository(db)
	courierRepo := repository.NewCourierRepository(db)
	myWalletRepo := repository.NewWalletTransactionHistoryRepository(db)

	pouc := usecase.ProductOrderUsecaseConfig{
		ProductOrderRepository:              productOrderRepo,
		ProductVariantCombinationRepository: productVariantCombinationRepo,
		AccountRepository:                   accountRepo,
		AccountAddressRepository:            accountAddressRepo,
		CourierRepository:                   courierRepo,
	}
	productRepo := repository.NewProductRepository(db)

	auc := usecase.AccountUsecaseConfig{
		AccountRepository:   accountRepo,
		UsedEmailRepository: usedEmailRepo,
		ProductRepository:   productRepo,
	}
	puc := usecase.ProductUsecaseConfig{
		ProductRepository: productRepo,
	}

	wuc := usecase.MyWalletTransactionUsecaseConfig{
		WalletTransactionRepo: myWalletRepo,
		AccountRepository:     accountRepo,
	}

	productUsecase := usecase.NewProductUsecase(puc)
	accountUsecase := usecase.NewAccountUsecase(auc)
	productOrderUsecase := usecase.NewProductOrderUsecase(pouc)
	walletTransactionUsecase := usecase.NewMyWalletTransactionUsecase(wuc)

	accountHandler := handler.NewAccountHandler(accountUsecase, walletTransactionUsecase)
	productOrderHandler := handler.NewProductOrderHandler(productOrderUsecase)
	productHandler := handler.NewProductHandler(productUsecase)

	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AddAllowHeaders("authorization")

	gin.Use(cors.New(configCors))

	router.NewPingRouter(gin)
	router.NewAccountRouter(accountHandler, gin)
	router.NewAuthRouter(accountHandler, gin)
	router_seller.NewProductOrderRouter(productOrderHandler, gin)
	router.NewProductRouter(productHandler, gin)
	router.NewProductOrderRouter(productOrderHandler, gin)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.GetEnv("PORT")),
		Handler: gin,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
