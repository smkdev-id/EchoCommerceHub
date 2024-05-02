package main

import (
	"smkdevid/echocommercehub/internal/configs"
	postgresql "smkdevid/echocommercehub/internal/databases/postgresql"
	"smkdevid/echocommercehub/internal/services/promotions"
	"smkdevid/echocommercehub/internal/transports/delivery"

	"github.com/labstack/echo/v4"
)

func main() {

	// Initialize Environment Variables
	configs.LoadViperEnv()

	// Initialize PostgreSQL Conn
	db := configs.InitDatabase()

	e := echo.New()

	// Apps Architect
	PromotionRepo := postgresql.NewPromotionRepository(db)

	PromoService := promotions.NewPromotionService(PromotionRepo)

	delivery.PromotionRoute(e, PromoService)

	e.Logger.Fatal(e.Start(":8080"))
}
