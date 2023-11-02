package middleware

import (
	"net/http"
	"os"

	"strings"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/gin-gonic/gin"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("ENV") == "testing" {
			c.Next()
			return
		}

		header := c.GetHeader("Authorization")
		s := strings.Split(header, " ")
		if len(s) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dtogeneral.JSONResponse{Message: "unauthorized"})
			return
		}
		s[0] = strings.ToLower(s[0])
		if s[0] != "bearer" && s[0] != "token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dtogeneral.JSONResponse{Message: "unauthorized"})
			return
		}

		token, err := util.ValidateToken(s[1])

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dtogeneral.JSONResponse{Message: "invalid token"})
			return
		}

		claims, ok := token.Claims.(*dtogeneral.ClaimsJWT)

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dtogeneral.JSONResponse{Message: "unauthorized"})
			return
		}
		c.Set("userId", claims.UserId)
		c.Next()
	}
}
