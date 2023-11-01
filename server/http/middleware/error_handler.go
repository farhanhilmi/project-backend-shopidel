package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"

	dtogeneral "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/general"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		log.Println("ERR", err)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				c.AbortWithStatusJSON(http.StatusBadGateway, dtogeneral.ErrResponse{Error: "request timeout"})
			}

			errMap, ok := err.Err.(*util.CustomError)

			if !ok {
				c.AbortWithStatusJSON(http.StatusInternalServerError, dtogeneral.ErrResponse{Error: "Internal Server Error"})
			}

			switch errMap.Code {
			case util.BadRequest:
				c.AbortWithStatusJSON(http.StatusBadRequest, dtogeneral.ErrResponse{Error: errMap.Message})
			case util.Unauthorized:
				c.AbortWithStatusJSON(http.StatusUnauthorized, dtogeneral.ErrResponse{Error: errMap.Message})
			case util.NotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, dtogeneral.ErrResponse{Error: errMap.Message})

			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, dtogeneral.ErrResponse{Error: "Internal Server Error"})
			}
			c.Abort()
		}
	}
}
