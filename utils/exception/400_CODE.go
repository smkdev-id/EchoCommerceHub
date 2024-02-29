package exception

import "fmt"

type NotFoundError struct {
	Message string
	ID      uint
}

type PromotionIDNotFoundError struct {
	Message     string
	PromotionID string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %d", e.Message, e.ID)
}

func (e *PromotionIDNotFoundError) Error() string {
	return fmt.Sprintf("%s with Promotion ID %s", e.Message, e.PromotionID)
}
