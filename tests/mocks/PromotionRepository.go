package mocks

import (
	schema "smkdevid/echocommercehub/internal/models/schema"

	"time"

	"github.com/stretchr/testify/mock"
)

type MockPromotionRepository struct {
	mock.Mock
}

func (m *MockPromotionRepository) CreatePromotion(promo schema.Promotion) (schema.Promotion, error) {
	promo.ID = 1
	promo.CreatedAt = time.Now()
	promo.UpdatedAt = time.Now()

	args := m.Called(promo)
	return args.Get(0).(schema.Promotion), args.Error(1)
}

func (m *MockPromotionRepository) GetAllPromotions() ([]schema.Promotion, error) {
	args := m.Called()
	return args.Get(0).([]schema.Promotion), args.Error(1)
}

func (m *MockPromotionRepository) GetPromotionbyPromotionID(promotionID string) (schema.Promotion, error) {
	args := m.Called(promotionID)
	return args.Get(0).(schema.Promotion), args.Error(1)
}

func (m *MockPromotionRepository) UpdatePromotionbyPromotionID(promo schema.Promotion) (schema.Promotion, error) {
	args := m.Called(promo)
	return args.Get(0).(schema.Promotion), args.Error(1)
}

func (m *MockPromotionRepository) DeletePromotionbyPromotionID(promotionID string) error {
	args := m.Called(promotionID)
	return args.Error(0)
}
