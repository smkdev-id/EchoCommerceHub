package configs

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadViperEnv() {
	viper.AddConfigPath("/")
	viper.SetConfigFile("env.yaml")
	viper.AutomaticEnv()

	if EnvException := viper.ReadInConfig(); EnvException != nil {
		panic(EnvException)
	}
}

func InitDatabase() *gorm.DB {
	dbUser := viper.GetString("DATABASE.USER")
	dbPass := viper.GetString("DATABASE.PASS")
	dbHost := viper.GetString("DATABASE.HOST")
	dbPort := viper.GetString("DATABASE.PORT")
	dbName := viper.GetString("DATABASE.NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbName, dbPort)

	// Create a new database connection with proper error handling
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
		// fmt.Errorf("failed to connect to database: %w", err)
	}

	// Optionally run migrations using `migrations(DB)`
	// migrations(DB)

	// Set logging level (adjust as needed)

	return DB
}
