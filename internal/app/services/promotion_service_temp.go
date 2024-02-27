package services

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/models"
	"gorm.io/gorm"
)

func TempPSQLCreatePromotionData(c echo.Context) error {
	var promo models.Promotion

	if err := c.Bind(&promo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion data")
	}

	// Duplicate check (assuming ProductID and start/end date are unique)
	if result := db.Where("promotion_id = ? AND start_date <= ? AND end_date >= ?", promo.PromotionID, promo.StartDate, promo.EndDate).First(&models.Promotion{}); result.Error == nil {
		return echo.NewHTTPError(http.StatusConflict, "Duplicate promotion found")
	}

	if err := db.Create(&promo).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create promotion")
	}
	return c.JSON(http.StatusCreated, promo)
}

func TempPSQLGetAllPromotionData(c echo.Context) error {
	var promotions []models.Promotion

	// Find all promotions from the database
	if err := db.Debug().Find(&promotions).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve promotions"+err.Error())
	}

	// Close the database connection (optional, preferred for larger applications)
	// defer db.Close()

	// Return the retrieved promotions in JSON format
	return c.JSON(http.StatusOK, promotions)

}

func TempPSQLGetPromotionByID(c echo.Context) error {
	promotion_id, err := strconv.Atoi(c.Param("promotion_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion ID")
	}

	var promo models.Promotion
	if err := db.First(&promo, promotion_id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Promotion not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get promotion")
	}
	return c.JSON(http.StatusOK, promo)
}

func TempPSQLUpdatePromotionByID(c echo.Context) error {
	promotion_id, err := strconv.Atoi(c.Param("promotion_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion ID")
	}

	var promo models.Promotion
	if err := db.First(&promo, promotion_id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Promotion not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get promotion")
	}

	if err := c.Bind(&promo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion data")
	}

	// Duplicate check (excluding current ID)
	if result := db.Where("promotion_id = ? AND start_date <= ? AND end_date >= ? AND id != ?", promo.PromotionID, promo.StartDate, promo.EndDate, promo.PromotionID).First(&models.Promotion{}); result.Error == nil {
		return echo.NewHTTPError(http.StatusConflict, "Duplicate promotion found")
	}

	if err := db.Save(&promo).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update promotion")
	}
	return c.JSON(http.StatusOK, promo)
}

func TempPSQLDeletePromotionByID(c echo.Context) error {
	promotion_id, err := strconv.Atoi(c.Param("promotion_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion ID")
	}

	if err := db.Delete(&models.Promotion{}, promotion_id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Promotion not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete promotion")
	}
	return c.JSON(http.StatusNoContent, "Promotion Data deleted sucessfully")
}
