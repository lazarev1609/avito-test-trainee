package service

import (
	"avito-test-trainee/internal/models"
	"avito-test-trainee/internal/repository"
	"context"
)

type BannerService struct {
	bannerRepo repository.BannerRepository
}

func NewBannerService(bannerRepo repository.BannerRepository) *BannerService {
	return &BannerService{
		bannerRepo: bannerRepo,
	}
}

func (s *BannerService) GetUserBanner(ctx context.Context, tagID, featureID int, useLastRevision bool) (*models.Banner, error) {
	return s.bannerRepo.GetBanner(ctx, tagID, featureID, useLastRevision)
}

func (s *BannerService) GetAllBanners(ctx context.Context, featureID, tagID, limit, offset int) ([]*models.Banner, error) {
	return s.bannerRepo.GetAllBanners(ctx, featureID, tagID, limit, offset)
}

func (s *BannerService) CreateBanner(ctx context.Context, banner *models.Banner) error {
	return s.bannerRepo.CreateBanner(ctx, banner)
}

func (s *BannerService) UpdateBanner(ctx context.Context, id int, banner *models.Banner) error {
	return s.bannerRepo.UpdateBanner(ctx, id, banner)
}

func (s *BannerService) DeleteBanner(ctx context.Context, id int) error {
	return s.bannerRepo.DeleteBanner(ctx, id)
}
