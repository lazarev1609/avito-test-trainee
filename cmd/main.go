package main

import (
	"avito-test-trainee/internal/handlers"
	"avito-test-trainee/internal/repository"
	"avito-test-trainee/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Создание маршрутизатора Gin
	router := gin.Default()

	// Инициализация слоя репозитория
	bannerRepo := repository.NewInMemoryBannerRepository()

	// Инициализация слоя сервиса
	bannerService := service.NewBannerService(bannerRepo)

	// Инициализация маршрутов
	handlers.InitRoutes(router, *bannerService)

	// Запуск сервера
	router.Run(":8080")
}
