package main

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/config"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/database"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := database.GetInstance()

	gin := gin.Default()

	server.Start(gin, db)
}
