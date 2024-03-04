package models

import (
	"time"

	"gorm.io/gorm"
)

type Promotion struct {
	gorm.Model
	ID                 uint           `gorm:"primarykey"`
	PromotionID        string         `gorm:"column:promotion_id" json:"promotion_id"`
	PromotionName      string         `gorm:"not null" json:"promotion_name"`
	DiscountType       string         `gorm:"not null" json:"discount_type"`
	DiscountValue      float64        `gorm:"not null" json:"discount_value"`
	PromotionStartDate time.Time      `gorm:"not null" json:"promotion_start_date"`
	PromotionEndDate   time.Time      `gorm:"not null" json:"promotion_end_date"`
	CreatedAt          time.Time      `gorm:"autoCreatedTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

func (Promotion) TableName() string {
	return "promotion_table"
}
