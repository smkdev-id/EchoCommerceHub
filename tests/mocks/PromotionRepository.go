package mocks

import (
	"time"

	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type MockPromotionRepository struct {
	mock.Mock
}

func (m *MockPromotionRepository) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	promo.ID = 1
	promo.CreatedAt = time.Now()
	promo.UpdatedAt = time.Now()

	args := m.Called(promo)
	return args.Get(0).(models.Promotion), args.Error(1)
}

func (m *MockPromotionRepository) GetAllPromotions() ([]models.Promotion, error) {
	args := m.Called()
	return args.Get(0).([]models.Promotion), args.Error(1)
}

func (m *MockPromotionRepository) GetPromotionbyPromotionID(promotionID string) (models.Promotion, error) {
	args := m.Called(promotionID)
	return args.Get(0).(models.Promotion), args.Error(1)
}

func (m *MockPromotionRepository) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	args := m.Called(promo)
	return args.Get(0).(models.Promotion), args.Error(1)
}

func (m *MockPromotionRepository) DeletePromotionbyPromotionID(promotionID string) error {
	args := m.Called(promotionID)
	return args.Error(0)
}
