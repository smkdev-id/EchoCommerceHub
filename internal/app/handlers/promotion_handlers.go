package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/models"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/services"
	"github.com/smkdev-id/promotion_tracking_dashboard/utils/exception"
)

func PSQLCreatePromotionData(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var promo models.Promotion
		if err := c.Bind(&promo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion data")
		}

		createdPromo, err := PromoService.CreatePromotion(promo)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create promotion")
		}

		return c.JSON(http.StatusCreated, createdPromo)
	}
}

func PSQLGetAllPromotionData(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		promotions, err := PromoService.GetAllPromotions()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve promotions: "+err.Error())
		}
		return c.JSON(http.StatusOK, promotions)
	}
}

func PSQLGetPromotionByID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion ID")
		}

		promo, err := PromoService.GetPromotionByID(uint(id))
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

func PSQLGetPromotionByPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		promotionID := c.Param("promotion_id")

		promo, err := PromoService.GetPromotionByPromotionID(promotionID)
		if err != nil {

			// ! Update the exception with the custom one. For now leave it there.
			if e, ok := err.(*exception.PromotionIDNotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get promotion")
		}

		return c.JSON(http.StatusOK, promo)
	}
}

func PSQLUpdatePromotionByPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		promotionID := c.Param("promotion_id")

		// TODO: check if promotion_id is exist
		promo, err := PromoService.GetPromotionByPromotionID(promotionID)
		if err != nil {
			if e, ok := err.(*exception.PromotionIDNotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get promotion")
		}

		if err := c.Bind(&promo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion data")
		}

		// Update promotion
		updatedPromo, err := PromoService.UpdatePromotion(promo)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update promotion")
		}

		return c.JSON(http.StatusOK, updatedPromo)
	}
}

func PSQLDeletePromotionByPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		promotionID := c.Param("promotion_id")

		if err := PromoService.DeletePromotionByPromotionID(promotionID); err != nil {
			if e, ok := err.(*exception.NotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete promotion")
		}
		return c.JSON(http.StatusNoContent, "Promotion Data deleted successfully") // 204
	}
}