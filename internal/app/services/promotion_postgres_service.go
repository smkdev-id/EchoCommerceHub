package services

import (
	"errors"

	"gorm.io/gorm"

	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/models"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/repositories"
	"github.com/smkdev-id/promotion_tracking_dashboard/utils/exception"
)

// PromotionService provides promotion-related services
type PromotionService interface {
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	GetAllPromotions() ([]models.Promotion, error)
	// GetPromotionbyID(ID uint) (models.Promotion, error)
	GetPromotionbyPromotionID(promotionID string) (models.Promotion, error)
	UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error)
	DeletePromotionbyPromotionID(promotionID string) error
}

type PromotionServiceImpl struct {
	PromotionRepo repositories.PromotionRepository
}

// NewPromotionService creates a new instance of PromotionService
func NewPromotionService(PromotionRepo repositories.PromotionRepository) *PromotionServiceImpl {
	return &PromotionServiceImpl{
		PromotionRepo: PromotionRepo,
	}
}

// CreatePromotion creates a new promotion
func (s *PromotionServiceImpl) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	return s.PromotionRepo.CreatePromotion(promo)
}

// GetAllPromotions that already recorded on database
func (s *PromotionServiceImpl) GetAllPromotions() ([]models.Promotion, error) {
	return s.PromotionRepo.GetAllPromotions()
}

// GetPromotionByPromotionID will throw data based on promotionID request
func (s *PromotionServiceImpl) GetPromotionbyPromotionID(promotionID string) (models.Promotion, error) {
	promo, err := s.PromotionRepo.GetPromotionbyPromotionID(promotionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Promotion{}, &exception.PromotionIDNotFoundError{
				Message:     "Promotion Not Found",
				PromotionID: promotionID,
			}
		}
		return models.Promotion{}, err
	}
	return promo, nil
}

// GetPromotionByID will throw data based on ID request
// func (s *PromotionServiceImpl) GetPromotionbyID(ID uint) (models.Promotion, error) {
// 	promo, err := s.PromotionRepo.GetPromotionbyID(uint(ID))
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return models.Promotion{}, &exception.NotFoundError{
// 				Message: "Promotion Not Found",
// 				ID:      ID,
// 			}
// 		}
// 		return models.Promotion{}, err
// 	}
// 	return promo, nil
// }

// UpdatePromotion will update data based on promotionID request
func (s *PromotionServiceImpl) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	// Perform duplicate check and update promotion
	updatePromo, err := s.PromotionRepo.UpdatePromotionbyPromotionID(promo)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Promotion{}, &exception.PromotionIDNotFoundError{
				Message:     "Duplicate Promotion Found",
				PromotionID: promo.PromotionID,
			}
		}
		return models.Promotion{}, err
	}
	return updatePromo, nil
}

// DeletePromotionByPromotionID will delete data based on promotionID request
func (s *PromotionServiceImpl) DeletePromotionbyPromotionID(promotionID string) error {
	return s.PromotionRepo.DeletePromotionbyPromotionID(promotionID)
}
