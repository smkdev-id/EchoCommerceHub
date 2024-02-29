package configs

import (
	"context"

	"github.com/jomei/notionapi"
	"github.com/spf13/viper"
)

var (
	NotionToken  = viper.GetString("NOTION.AUTH")
	NotionPageID = viper.GetString("NOTION_ID")
)

func NotionInit() (*notionapi.Page, error) {
	notion_client := notionapi.NewClient(notionapi.Token(NotionToken))

	promotion_report, err := notion_client.Page.Get(context.Background(), notionapi.PageID(NotionPageID))

	if err != nil {
		panic(err)
	}

	return promotion_report, nil
}
