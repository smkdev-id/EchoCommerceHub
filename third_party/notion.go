package thirdparty

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func NotionAuth(c echo.Context) {
}

func NotionDataSync() {
	viper.AddConfigPath(".env/")
}

func NotionUpdateData() {

}

func NotionDeleteData() {

}

func NotionPostData() {

}
