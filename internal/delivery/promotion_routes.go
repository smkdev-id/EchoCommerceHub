package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/handlers"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/services"
)

func HelloServer(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func PromotionRoute(e *echo.Echo, PromoService services.PromotionService) {

	e.GET("/", HelloServer)
	e.GET("/promotions", handlers.PSQLGetAllPromotionData(PromoService))
	e.GET("/promotions/:promotion_id", handlers.PSQLGetPromotionByID(PromoService))
	e.POST("/promotions", handlers.PSQLCreatePromotionData(PromoService))
	e.PUT("/promotions/:promotion_id", handlers.PSQLUpdatePromotionByID(PromoService))
	e.DELETE("/promotions/:promotion_id", handlers.PSQLDeletePromotionByID(PromoService))
}
