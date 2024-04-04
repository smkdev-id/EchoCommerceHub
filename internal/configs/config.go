package configs

import (
	models "smkdevid/echocommercehub/internal/models/schema"

	"github.com/jomei/notionapi"
	"github.com/spf13/viper"
)

type NotionReportAutomationService interface {
	PromotionUpdateData(promotion []models.Promotion) (string, error)
}

type NotionServiceImpl struct {
	NotionAuth *notionapi.Client
}

func LoadViperEnv() {
	viper.AddConfigPath("/")
	viper.SetConfigFile("env.yaml")
	viper.AutomaticEnv()

	if EnvException := viper.ReadInConfig(); EnvException != nil {
		panic(EnvException)
	}
}
