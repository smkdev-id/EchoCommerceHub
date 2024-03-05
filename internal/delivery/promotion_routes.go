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
	e.GET("/getpromotion/:promotion_id", handlers.PSQLGetPromotionbyPromotionID(PromoService))
	e.POST("/createpromotion", handlers.PSQLCreatePromotionData(PromoService))
	e.PUT("/updatepromotion/:promotion_id", handlers.PSQLUpdatePromotionbyPromotionID(PromoService))
	e.DELETE("/deletepromotion/:promotion_id", handlers.PSQLDeletePromotionbyPromotionID(PromoService))
}
