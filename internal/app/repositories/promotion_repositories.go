package repositories

import (
	"time"

	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/models"
	"github.com/smkdev-id/promotion_tracking_dashboard/utils/exception"
	"gorm.io/gorm"
)

type PromotionRepository interface {
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	FindPromotionByDateRange(promotionID uint, startDate, endDate time.Time) (*models.Promotion, error)
	GetAllPromotions() ([]models.Promotion, error)
	GetPromotionByID(promotionID uint) (models.Promotion, error)
	UpdatePromotion(promo models.Promotion) (models.Promotion, error)
	DeletePromotionByID(promotionID uint) error
}

type PromotionRepositoryImpl struct {
	db *gorm.DB
}

// NewPromotionRepository creates a new instance of PromotionRepository
func NewPromotionRepository(db *gorm.DB) PromotionRepository {
	return &PromotionRepositoryImpl{
		db: db,
	}
}

// CreatePromotion creates a new promotion in the database
func (r *PromotionRepositoryImpl) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	err := r.db.Create(&promo).Error
	return promo, err
}

// FindPromotionByDateRange finds a promotion by date range in the database
func (r *PromotionRepositoryImpl) FindPromotionByDateRange(promotionID uint, startDate, endDate time.Time) (*models.Promotion, error) {
	var promo models.Promotion
	result := r.db.Where("promotion_id = ? AND start_date <= ? AND end_date >= ?", promotionID, startDate, endDate).First(&promo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &promo, nil
}

func (r *PromotionRepositoryImpl) GetAllPromotions() ([]models.Promotion, error) {
	var promotions []models.Promotion
	if err := r.db.Debug().Find(&promotions).Error; err != nil {
		return nil, err
	}
	return promotions, nil
}

func (r *PromotionRepositoryImpl) GetPromotionByID(promotionID uint) (models.Promotion, error) {
	var promo models.Promotion
	if err := r.db.First(&promo, promotionID).Error; err != nil {
		return models.Promotion{}, err
	}
	return promo, nil
}

func (r *PromotionRepositoryImpl) UpdatePromotion(promo models.Promotion) (models.Promotion, error) {
	// Assuming you have unique constraints on promotion_id and dates,
	// you can perform the duplicate check before updating
	var existingPromo models.Promotion
	if err := r.db.Where("promotion_id = ? AND start_date <= ? AND end_date >= ? AND id != ?", promo.PromotionID, promo.StartDate, promo.EndDate, promo.ID).First(&existingPromo).Error; err != nil {
		// Duplicate promotion found
		return models.Promotion{}, err
	}

	// Update the promotion
	if err := r.db.Save(&promo).Error; err != nil {
		return models.Promotion{}, err
	}
	return promo, nil
}

func (r *PromotionRepositoryImpl) DeletePromotionByID(promotionID uint) error {
	if err := r.db.Delete(&models.Promotion{}, promotionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &exception.NotFoundError{
				Message: "Promotion Not Found",
				ID:      promotionID,
			}
		}
		return err
	}
	return nil
}
