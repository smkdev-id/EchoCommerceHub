package mocks

import (
	schema "smkdevid/echocommercehub/internal/models/schema"

	"github.com/stretchr/testify/mock"
)

type MockPromotionService struct {
	mock.Mock
}

func (m *MockPromotionService) CreatePromotion(promo schema.Promotion) (schema.Promotion, error) {
	args := m.Called(promo)
	return args.Get(0).(schema.Promotion), args.Error(1)
}

func (m *MockPromotionService) GetAllPromotions() ([]schema.Promotion, error) {
	args := m.Called()
	return args.Get(0).([]schema.Promotion), args.Error(1)
}

func (m *MockPromotionService) GetPromotionbyPromotionID(promotionID string) (schema.Promotion, error) {
	args := m.Called(promotionID)
	if result := args.Get(0); result != nil {
		return result.(schema.Promotion), nil
	}
	return schema.Promotion{}, args.Error(1)
}

func (m *MockPromotionService) UpdatePromotionbyPromotionID(promo schema.Promotion) (schema.Promotion, error) {
	args := m.Called(promo)
	if result := args.Get(0); result != nil {
		return result.(schema.Promotion), nil
	}
	return schema.Promotion{}, args.Error(1)
}

func (m *MockPromotionService) DeletePromotionbyPromotionID(promotionID string) error {
	args := m.Called(promotionID)
	return args.Error(0)
}
