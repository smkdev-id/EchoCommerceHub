package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/smkdev-id/promotion_tracking_dashboard/services"
)

func HelloServer(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()

	e.GET("/", HelloServer)
	e.GET("/getpromotiondata", services.GetAllPromotionData)
	e.POST("/createpromotiondata", services.CreatePromotionData)
	e.GET("/getpromotiondata/:promotion_id", services.GetPromotionByID)
	e.PUT("/updatepromotiondata/:promotion_id", services.UpdatePromotionByID)
	e.DELETE("/deletepromotiondata/:promotion_id", services.DeletePromotionByID)

	e.Logger.Fatal(e.Start(":8080"))
}
