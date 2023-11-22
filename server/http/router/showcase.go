package router

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewShowcaseRouter(h *handler.ShowcaseHandler, gin *gin.Engine) *gin.Engine {
	showcase := gin.Group("/showcases")
	showcase.Use(middleware.AuthenticateJWT())

	showcase.GET("", h.GetShowcases)
	showcase.POST("", h.CreateShowcase)
	showcase.GET("/:showcaseId", h.GetShowcaseDetail)
	showcase.PUT("/:showcaseId", h.UpdateShowcase)
	showcase.DELETE("/:showcaseId", h.DeleteShowcase)

	return gin
}
