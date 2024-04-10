package handlers

import (
	"avito-test-trainee/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// BannerHandler представляет собой обработчик для эндпоинтов, связанных с баннерами
type BannerHandler struct {
	bannerService service.BannerService
}

// NewBannerHandler создает новый экземпляр BannerHandler
func NewBannerHandler(bannerService service.BannerService) *BannerHandler {
	return &BannerHandler{
		bannerService: bannerService,
	}
}

func (h *BannerHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/user_banner", h.GetUserBanner)
	router.GET("/banner", h.GetAllBanners)
	router.POST("/banner", h.CreateBanner)
	router.PATCH("/banner/:id", h.UpdateBanner)
	router.DELETE("/banner/:id", h.DeleteBanner)
}

func (h *BannerHandler) GetUserBanner(c *gin.Context) {
	tagID, _ := strconv.Atoi(c.Query("tag_id"))
	featureID, _ := strconv.Atoi(c.Query("feature_id"))
	useLastRevision, _ := strconv.ParseBool(c.Query("use_last_revision"))

	banner, err := h.bannerService.GetUserBanner(context.Background(), tagID, featureID, useLastRevision)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, banner)
}

func (h *BannerHandler) GetAllBanners(c *gin.Context) {
	featureID, _ := strconv.Atoi(c.Query("feature_id"))
	tagID, _ := strconv.Atoi(c.Query("tag_id"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	banners, err := h.bannerService.GetAllBanners(context.Background(), featureID, tagID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, banners)
}

func (h *BannerHandler) CreateBanner(c *gin.Context) {
	// Логика для создания нового баннера
}

func (h *BannerHandler) UpdateBanner(c *gin.Context) {
	// Логика для обновления содержимого баннера
}

func (h *BannerHandler) DeleteBanner(c *gin.Context) {
	// Логика для удаления баннера
}
