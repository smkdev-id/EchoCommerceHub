package configs

// import (
// 	"context"
// 	"fmt"

// 	"github.com/jomei/notionapi"
// 	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/models"
// 	"github.com/spf13/viper"
// )

// type NotionReportAutomationService interface {
// 	PromotionUpdateData(promotion []models.Promotion) string
// }

// type NotionServiceImpl struct {
// 	NotionAuth *notionapi.Client
// }

// func NotionInit() NotionReportAutomationService {
// 	notion_client := notionapi.NewClient(notionapi.Token(viper.GetString("NOTION.AUTH")))

// 	promotion_report, err := notion_client.Page.Get(context.Background(), notionapi.PageID(viper.GetString("NOTION_ID")))

// 	if err != nil{
// 		return nil, fmt.Errorf("Failed to Connet to Notion Page: %w", err)
// 	}

// 	return promotion_report,
// }
