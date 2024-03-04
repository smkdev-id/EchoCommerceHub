package mocks

import (
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type MockPromotionService struct {
	mock.Mock
}

func (m *MockPromotionService) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	args := m.Called(promo)
	return args.Get(0).(models.Promotion), args.Error(1)
}

func (m *MockPromotionService) GetAllPromotions() ([]models.Promotion, error) {
	args := m.Called()
	return args.Get(0).([]models.Promotion), args.Error(1)
}

func (m *MockPromotionService) GetPromotionbyPromotionID(promotionID string) (models.Promotion, error) {
	args := m.Called(promotionID)
	if result := args.Get(0); result != nil {
		return result.(models.Promotion), nil
	}
	return models.Promotion{}, args.Error(1)
}

func (m *MockPromotionService) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	args := m.Called(promo)
	if result := args.Get(0); result != nil {
		return result.(models.Promotion), nil
	}
	return models.Promotion{}, args.Error(1)
}

func (m *MockPromotionService) DeletePromotionbyPromotionID(promotionID string) error {
	args := m.Called(promotionID)
	return args.Error(0)
}
