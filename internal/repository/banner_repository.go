package repository

import (
	"avito-test-trainee/internal/models"
	"context"
)

type BannerRepository interface {
	GetBanner(ctx context.Context, tagID, featureID int, useLastRevision bool) (*models.Banner, error)
	GetAllBanners(ctx context.Context, featureID, tagID, limit, offset int) ([]*models.Banner, error)
	CreateBanner(ctx context.Context, banner *models.Banner) error
	UpdateBanner(ctx context.Context, id int, banner *models.Banner) error
	DeleteBanner(ctx context.Context, id int) error
}

type InMemoryBannerRepository struct {
	banners []*models.Banner
}

func NewInMemoryBannerRepository() *InMemoryBannerRepository {
	return &InMemoryBannerRepository{
		banners: make([]*models.Banner, 0),
	}
}

func (r *InMemoryBannerRepository) GetBanner(ctx context.Context, tagID, featureID int, useLastRevision bool) (*models.Banner, error) {
	// Логика для получения баннера из хранилища
	return nil, nil
}

func (r *InMemoryBannerRepository) GetAllBanners(ctx context.Context, featureID, tagID, limit, offset int) ([]*models.Banner, error) {
	// Логика для получения всех баннеров с учетом переданных параметров
	return make([]*models.Banner, 0), nil
}

func (r *InMemoryBannerRepository) CreateBanner(ctx context.Context, banner *models.Banner) error {
	// Логика для создания нового баннера
	return nil
}

func (r *InMemoryBannerRepository) UpdateBanner(ctx context.Context, id int, banner *models.Banner) error {
	// Логика для обновления баннера
	return nil
}

func (r *InMemoryBannerRepository) DeleteBanner(ctx context.Context, id int) error {
	// Логика для удаления баннера
	return nil
}
