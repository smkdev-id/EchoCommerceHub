package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/models"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/services"
	"github.com/smkdev-id/promotion_tracking_dashboard/utils/exception"
)

func PSQLCreatePromotionData(promoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var promo models.Promotion
		if err := c.Bind(&promo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion data")
		}

		createdPromo, err := promoService.CreatePromotion(promo)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create promotion")
		}

		return c.JSON(http.StatusCreated, createdPromo)
	}
}

func PSQLGetAllPromotionData(promoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		promotions, err := promoService.GetAllPromotions()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve promotions: "+err.Error())
		}
		return c.JSON(http.StatusOK, promotions)
	}
}

func PSQLGetPromotionByID(promoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		promotionID, err := strconv.Atoi(c.Param("promotion_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion ID")
		}

		promo, err := promoService.GetPromotionByID(uint(promotionID))
		if err != nil {

			// ! Update the exception with the custom one. For now leave it there.
			if e, ok := err.(*exception.NotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get promotion")
		}

		return c.JSON(http.StatusOK, promo)
	}
}

func PSQLUpdatePromotionByID(promoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		promotionID, err := strconv.Atoi(c.Param("promotion_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion ID")
		}

		promo, err := promoService.GetPromotionByID(uint(promotionID))
		if err != nil {
			if e, ok := err.(*exception.NotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get promotion")
		}

		if err := c.Bind(&promo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion data")
		}

		// Update promotion
		updatedPromo, err := promoService.UpdatePromotionByID(promo)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update promotion")
		}

		return c.JSON(http.StatusOK, updatedPromo)
	}
}

func PSQLDeletePromotionByID(promoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		promotionID, err := strconv.Atoi(c.Param("promotion_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion ID")
		}

		if err := promoService.DeletePromotionByID(uint(promotionID)); err != nil {
			if e, ok := err.(*exception.NotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete promotion")
		}
		return c.JSON(http.StatusNoContent, "Promotion Data deleted successfully") // 204
	}
}
