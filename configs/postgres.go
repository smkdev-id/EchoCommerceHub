package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PromotionData *gorm.DB

func InitDatabase() *gorm.DB {
	dbUser := viper.GetString("DATABASE.USER")
	dbPass := viper.GetString("DATABASE.PASS")
	dbHost := viper.GetString("DATABASE.HOST")
	dbPort := viper.GetString("DATABASE.PORT")
	dbName := viper.GetString("DATABASE.NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbName, dbPort)

	var errDB error
	DB, errDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatal("Failed to Connect Database")
	}

	// migrations(DB)

	return DB
}
