package database

import (
	"fmt"
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/config"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
	if db == nil {
		db = connect()
	}

	return db
}

func connect() *gorm.DB {
	host := config.GetEnv("DB_HOST")
	port := config.GetEnv("DB_PORT")
	user := config.GetEnv("DB_USER")
	pass := config.GetEnv("DB_PASS")
	dbName := config.GetEnv("DB_NAME")
	sslMode := config.GetEnv("DB_SSL_MODE")

	portDB, _ := strconv.Atoi(port)

	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, user, pass, dbName, portDB, sslMode)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})
	if err != nil {
		panic(err)
	}

	return db
}

func MigrateDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Accounts{})
	if err != nil {
		return err
	}

	db.Exec(`
	ALTER TABLE accounts ALTER COLUMN wallet_number
		SET DEFAULT TO_CHAR(nextval('id'::regclass),'"420"fm0000000000');
	`)
	return nil
}
