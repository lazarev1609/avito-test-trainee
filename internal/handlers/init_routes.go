package handlers

import (
	"avito-test-trainee/internal/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, bannerService service.BannerService) {
	bannerHandler := NewBannerHandler(bannerService)
	bannerHandler.RegisterRoutes(router)
}
