package delivery

import (
	"net/http"

	"smkdevid/echocommercehub/internal/app/handlers"
	"smkdevid/echocommercehub/internal/services/promotions"

	"github.com/labstack/echo/v4"
)

func HelloServer(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func PromotionRoute(e *echo.Echo, PromoService promotions.PromotionService) {

	e.GET("/", HelloServer)
	e.GET("/promotions", handlers.PSQLGetAllPromotionData(PromoService))
	e.GET("/getpromotion/:promotion_id", handlers.PSQLGetPromotionbyPromotionID(PromoService))
	e.POST("/createpromotion", handlers.PSQLCreatePromotionData(PromoService))
	e.PUT("/updatepromotion/:promotion_id", handlers.PSQLUpdatePromotionbyPromotionID(PromoService))
	e.DELETE("/deletepromotion/:promotion_id", handlers.PSQLDeletePromotionbyPromotionID(PromoService))
}
