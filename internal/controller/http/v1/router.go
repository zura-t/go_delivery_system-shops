package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zura-t/go_delivery_system-shops/internal/usecase"
	"github.com/zura-t/go_delivery_system-shops/pkg/logger"
)

func NewRouter(handler *gin.Engine, logger logger.Interface, shopUsecase *usecase.ShopUseCase) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	h := handler.Group("/v1")
	handler.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	{
		newShopRoutes(h, shopUsecase, logger)
	}
}