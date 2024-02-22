package models

import (
	"time"

	"gorm.io/gorm"
)

type Promotion struct {
	gorm.Model
	PromotionID   uint      `gorm:"not null" json:"promotion_id"`
	Name          string    `gorm:"not null" json:"name"`
	DiscountType  string    `gorm:"not null" json:"discount_type"`
	DiscountValue float64   `gorm:"not null" json:"discount_value"`
	StartDate     time.Time `gorm:"not null" json:"start_date"`
	EndDate       time.Time `gorm:"not null" json:"end_date"`
}
