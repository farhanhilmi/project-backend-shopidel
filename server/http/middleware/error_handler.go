package middleware

import (
	"context"
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/util"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				c.AbortWithStatusJSON(http.StatusBadGateway, dto.ErrResponse{Error: "request timeout"})
			}

			errMap, ok := err.Err.(*util.CustomError)

			if !ok {
				c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrResponse{Error: "Internal Server Error"})
			}

			switch errMap.Code {
			case util.BadRequest:
				c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrResponse{Error: errMap.Message})
			case util.Unauthorized:
				c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrResponse{Error: errMap.Message})
			case util.NotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, dto.ErrResponse{Error: errMap.Message})

			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrResponse{Error: "Internal Server Error"})
			}
			c.Abort()
		}
	}
}
