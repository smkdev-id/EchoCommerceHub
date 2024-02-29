package services

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/models"
)

// Our Database
var promotions []models.Promotion

func CreatePromotionData(c echo.Context) error {
	var promo models.Promotion

	// Check Invalid Data based on Promotion Struct
	if err := c.Bind(&promo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Promotion Data")
	}

	// TODO: Check Duplicate Data interatively
	for _, p := range promotions {
		if p.ID == promo.ID && ((p.PromotionStartDate.Equal(promo.PromotionStartDate) || p.PromotionStartDate.Before(promo.PromotionStartDate)) && (p.PromotionEndDate.Equal(promo.PromotionEndDate) || p.PromotionEndDate.After(promo.PromotionEndDate))) {
			return echo.NewHTTPError(http.StatusConflict, "Duplicate promotion found")
		}
	}

	// Append Data to Database
	promotions = append(promotions, promo)

	// TODO: Sort the data based on start_date ASC
	sort.Slice(promotions, func(i, j int) bool {
		return promotions[i].PromotionStartDate.Before(promotions[j].PromotionStartDate)
	})

	// Return Data already inputted/created
	return c.JSON(http.StatusCreated, promo)
}

// Throw all recorded Data from Promotion
func GetAllPromotionData(c echo.Context) error {
	return c.JSON(http.StatusOK, promotions)
}

func GetPromotionByID(c echo.Context) error {
	promotion_id, err := strconv.Atoi(c.Param("promotion_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion ID")
	}

	// TODO: Iterate over the promotions slice to find the desired promotion
	for _, promo := range promotions {
		if int(promo.ID) == promotion_id {
			return c.JSON(http.StatusOK, promo)
		}
	}

	// If promotion with given ID is not found, return an error
	return echo.NewHTTPError(http.StatusNotFound, "Promotion not found")
}

func UpdatePromotionByID(c echo.Context) error {
	promotion_id, err := strconv.Atoi(c.Param("promotion_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion ID")
	}

	// Find the index of the promotion in the promotions slice
	index := -1
	for i, promo := range promotions {
		if int(promo.ID) == promotion_id {
			index = i
			break
		}
	}

	if index == -1 {
		return echo.NewHTTPError(http.StatusNotFound, "Promotion not found")
	}

	// Retrieve the existing promotion
	promo := promotions[index]

	// Update the promotion with the new data
	if err := c.Bind(&promo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion data")
	}

	// Duplicate check (excluding current ID)
	for _, p := range promotions {
		if p.ID == promo.ID && p.ID != promo.ID && ((p.PromotionStartDate.Before(promo.PromotionEndDate) || p.PromotionStartDate.Equal(promo.PromotionEndDate)) && (p.PromotionEndDate.After(promo.PromotionStartDate) || p.PromotionEndDate.Equal(promo.PromotionStartDate))) {
			return echo.NewHTTPError(http.StatusConflict, "Duplicate promotion found")
		}
	}

	// Update the promotion in the slice
	promotions[index] = promo

	return c.JSON(http.StatusOK, promo)
}

func DeletePromotionByID(c echo.Context) error {
	promotion_id, err := strconv.Atoi(c.Param("promotion_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion ID")
	}

	deletedIndex := -1
	for i, p := range promotions {
		if int(p.ID) == promotion_id {
			deletedIndex = i
			break
		}
	}

	if deletedIndex == -1 {
		return echo.NewHTTPError(http.StatusNotFound, "Promotion not found")
	}

	deletedPromotion := promotions[deletedIndex]
	promotions = append(promotions[:deletedIndex], promotions[deletedIndex+1:]...)

	return c.JSON(http.StatusNoContent, deletedPromotion)
}
