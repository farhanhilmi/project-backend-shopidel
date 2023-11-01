package middleware

import (
	"net/http"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	"github.com/gin-gonic/gin"
)

func CheckContentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.GetHeader("Content-Type")
		if contentType != "application/json" {
			c.AbortWithStatusJSON(http.StatusBadRequest, dtogeneral.ErrResponse{Error: "Invalid content type, expecting application/json"})
			return
		}
		c.Next()
	}
}
