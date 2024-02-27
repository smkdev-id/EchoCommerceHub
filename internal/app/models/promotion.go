package models

import (
	"time"

	"gorm.io/gorm"
)

// type Promotion struct {
// 	gorm.Model
// 	PromotionID   uint      `gorm:"primarykey;not null" json:"promotion_id"`
// 	Name          string    `gorm:"not null" json:"name"`
// 	DiscountType  string    `gorm:"not null" json:"discount_type"`
// 	DiscountValue float64   `gorm:"not null" json:"discount_value"`
// 	StartDate     time.Time `gorm:"not null" json:"start_date"`
// 	EndDate       time.Time `gorm:"not null" json:"end_date"`
// }

type Promotion struct {
	gorm.Model
	PromotionID uint   `gorm:"primarykey;column:promotion_id" json:"promotion_id"`
	Name        string `gorm:"not null" binding:"required" json:"name"`
	// DiscountType  DiscountType `gorm:"not null" binding:"required" json:"discount_type"`

	DiscountType  string    `gorm:"not null" json:"discount_type"`
	DiscountValue float64   `gorm:"not null" binding:"required" json:"discount_value"`
	StartDate     time.Time `gorm:"not null" binding:"required" json:"start_date"`
	EndDate       time.Time `gorm:"not null" binding:"required" json:"end_date"`
	// DeletedAt     time.Time `gorm:"softdelete"`
}

func (Promotion) TableName() string {
	return "promotion_table"
}

// // Disable Soft Deletes. Preventing deleted_at on schema
// func (Promotion) DeleteOption(db *gorm.DB) *gorm.DB {
// 	return db.Set("gorm:delete_option", "CASCADE")
// }
