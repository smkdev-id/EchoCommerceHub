package services

import (
	"errors"

	"gorm.io/gorm"

	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/models"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/repositories"
	"github.com/smkdev-id/promotion_tracking_dashboard/utils/exception"
)

var db *gorm.DB

// PromotionService provides promotion-related services
type PromotionService interface {
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	GetAllPromotions() ([]models.Promotion, error)
	GetPromotionByID(promotionID uint) (models.Promotion, error)
	UpdatePromotionByID(promo models.Promotion) (models.Promotion, error)
	DeletePromotionByID(promotionID uint) error
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
	// Duplicate check (assuming ProductID and start/end date are unique)
	existingPromo, err := s.PromotionRepo.FindPromotionByDateRange(promo.PromotionID, promo.StartDate, promo.EndDate)
	if err != nil {
		return models.Promotion{}, err
	}
	if existingPromo != nil {
		return models.Promotion{}, errors.New("duplicate promotion found")
	}

	return s.PromotionRepo.CreatePromotion(promo)
}

func (s *PromotionServiceImpl) GetAllPromotions() ([]models.Promotion, error) {
	return s.PromotionRepo.GetAllPromotions()
}

func (s *PromotionServiceImpl) GetPromotionByID(promotionID uint) (models.Promotion, error) {
	promo, err := s.PromotionRepo.GetPromotionByID(promotionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Promotion{}, &exception.NotFoundError{
				Message: "Promotion Not Found",
				ID:      promotionID,
			}
		}
		return models.Promotion{}, err
	}
	return promo, nil
}

func (s *PromotionServiceImpl) UpdatePromotionByID(promo models.Promotion) (models.Promotion, error) {
	// Perform duplicate check and update promotion
	updatePromo, err := s.PromotionRepo.UpdatePromotion(promo)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Promotion{}, &exception.NotFoundError{
				Message: "Duplicate Promotion Found",
				ID:      promo.PromotionID,
			}
		}
		return models.Promotion{}, err
	}
	return updatePromo, nil
}

func (s *PromotionServiceImpl) DeletePromotionByID(promotionID uint) error {
	return s.PromotionRepo.DeletePromotionByID(promotionID)
}
