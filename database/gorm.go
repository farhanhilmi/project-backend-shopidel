package database

import (
	"fmt"
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
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
		return nil, err
	}

	return db, nil
}
