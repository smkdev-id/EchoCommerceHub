package main

import (
	"github.com/labstack/echo/v4"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/repositories"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/services"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/configs"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/delivery"
)

func main() {

	// Initialize Environment Variables
	configs.LoadViperEnv()

	// Initialize PostgreSQL Conn
	db := configs.InitDatabase()

	e := echo.New()

	// Apps Architect
	PromotionRepo := repositories.NewPromotionRepository(db)

	PromoService := services.NewPromotionService(PromotionRepo)

	delivery.PromotionRoute(e, PromoService)

	e.Logger.Fatal(e.Start(":8080"))
}
