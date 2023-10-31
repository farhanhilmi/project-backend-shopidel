package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	"github.com/gin-gonic/gin"
)

func NewPingRouter(gin *gin.Engine) *gin.Engine {
	gin.GET("/ping", ping)
	return gin
}

func ping(c *gin.Context) {
	c.JSON(200, dto.JSONResponse{Data: "pong"})
}
