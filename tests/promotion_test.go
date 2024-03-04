package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/handlers"
	"github.com/smkdev-id/promotion_tracking_dashboard/internal/app/models"
	"github.com/smkdev-id/promotion_tracking_dashboard/tests/mocks"
	"github.com/smkdev-id/promotion_tracking_dashboard/utils/exception"
	"github.com/stretchr/testify/assert"
)

func TestCreatePromotion_Success(t *testing.T) {
	t.Run("Successful promotion creation", func(t *testing.T) {
		// Set up mocks
		mockRepo := new(mocks.MockPromotionRepository)
		mockService := new(mocks.MockPromotionService)

		// Replace with your actual model data
		expectedPromo := models.Promotion{
			PromotionID:        "cae8651b",
			PromotionName:      "Ramadhan Sale",
			DiscountType:       "percentage",
			DiscountValue:      10.5,
			PromotionStartDate: time.Now(),
			PromotionEndDate:   time.Now().Add(24 * time.Hour),
		}

		mockRepo.On("CreatePromotion", expectedPromo).Return(expectedPromo, nil)
		mockService.On("CreatePromotion", expectedPromo).Return(expectedPromo, nil)

		// Create handler with mocks
		e := echo.New()
		handler := handlers.PSQLCreatePromotionData(mockService)
		e.POST("/createpromotion", handler)

		payload, _ := json.Marshal(expectedPromo)

		// Create request
		req := httptest.NewRequest(http.MethodPost, "/createpromotion", bytes.NewReader(payload))
		rec := httptest.NewRecorder()

		// Execute request
		e.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusCreated, rec.Code)
		var actualPromo models.Promotion
		err := json.Unmarshal(rec.Body.Bytes(), &actualPromo)
		assert.NoError(t, err)
		assert.Equal(t, expectedPromo, actualPromo)

		// Assert mock interactions
		mockRepo.AssertExpectations(t)
		mockService.AssertExpectations(t)
	})

	t.Run("Error on promotion creation", func(t *testing.T) {
		// Set up mocks
		mockRepo := new(mocks.MockPromotionRepository)
		mockService := new(mocks.MockPromotionService)

		expectedPromo := models.Promotion{
			PromotionID:        "cae8651b",
			PromotionName:      "Ramadhan Sale",
			DiscountType:       "percentage",
			DiscountValue:      10.5,
			PromotionStartDate: time.Now(),
			PromotionEndDate:   time.Now().Add(24 * time.Hour),
		}

		expectedErr := errors.New("failed to create promotion")
		mockRepo.On("CreatePromotion", expectedPromo).Return(models.Promotion{}, expectedErr)
		mockService.On("CreatePromotion", expectedPromo).Return(models.Promotion{}, expectedErr)

		// Create handler with mocks
		e := echo.New()
		handler := handlers.PSQLCreatePromotionData(mockService)
		e.POST("/createpromotion", handler)

		payload, _ := json.Marshal(expectedPromo)

		// Create request
		req := httptest.NewRequest(http.MethodPost, "/createpromotion", bytes.NewReader(payload))
		rec := httptest.NewRecorder()

		// Execute request
		e.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body, err := io.ReadAll(rec.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(body), expectedErr.Error())

		// Assert mock interactions
		mockRepo.AssertExpectations(t)
		mockService.AssertExpectations(t)
	})
}

func TestPSQLGetAllPromotionData(t *testing.T) {
	t.Run("Successful retrieval of promotions", func(t *testing.T) {
		// Set up mocks
		mockRepo := new(mocks.MockPromotionRepository)
		mockService := new(mocks.MockPromotionService)

		promotionStartDate, _ := time.Parse(time.RFC3339, "2024-03-01T00:00:00+07:00")
		promotionEndDate, _ := time.Parse(time.RFC3339, "2024-03-31T23:59:59+07:00")
		createdAt, _ := time.Parse(time.RFC3339, "2024-02-29T15:29:18.047485+07:00")
		updateAt, _ := time.Parse(time.RFC3339, "2024-02-29T15:29:18.047485+07:00")

		expectedPromotions := []models.Promotion{
			{
				ID:                 1,
				PromotionID:        "cae8651b",
				PromotionName:      "Ramadhan Sale",
				DiscountType:       "Percentage",
				DiscountValue:      10.5,
				PromotionStartDate: promotionStartDate,
				PromotionEndDate:   promotionEndDate,
				CreatedAt:          createdAt,
				UpdatedAt:          updateAt,
				// DeletedAt:          nil, // Assuming soft deletion is disabled
			},
		}

		mockRepo.On("GetAllPromotions").Return(expectedPromotions, nil)
		mockService.On("GetAllPromotions").Return(expectedPromotions, nil)

		// Create handler with mocks
		e := echo.New()
		handler := handlers.PSQLGetAllPromotionData(mockService)
		e.GET("/promotions", handler)

		// Create request
		req := httptest.NewRequest(http.MethodGet, "/promotions", nil)
		rec := httptest.NewRecorder()

		// Execute request
		e.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualPromotions []models.Promotion
		err := json.Unmarshal(rec.Body.Bytes(), &actualPromotions)
		assert.NoError(t, err)
		assert.Equal(t, expectedPromotions, actualPromotions)

		// Assert mock interactions
		mockRepo.AssertExpectations(t)
		mockService.AssertExpectations(t)
	})

	t.Run("Error retrieving promotions", func(t *testing.T) {
		// Set up mocks
		mockRepo := new(mocks.MockPromotionRepository)
		mockService := new(mocks.MockPromotionService)

		expectedErr := errors.New("failed to get promotions")
		mockRepo.On("GetAllPromotions").Return(nil, expectedErr)
		mockService.On("GetAllPromotions").Return(nil, expectedErr)

		// Create handler with mocks
		e := echo.New()
		handler := handlers.PSQLGetAllPromotionData(mockService)
		e.GET("/promotions", handler)

		// Create request
		req := httptest.NewRequest(http.MethodGet, "/promotions", nil)
		rec := httptest.NewRecorder()

		// Execute request
		e.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		body, err := io.ReadAll(rec.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(body), expectedErr.Error())

		// Assert mock interactions
		mockRepo.AssertExpectations(t)
		mockService.AssertExpectations(t)
	})
}

func TestPSQLGetPromotionbyPromotionID(t *testing.T) {
	t.Run("Successful retrieval of promotion", func(t *testing.T) {
		// Set up mocks
		mockRepo := new(mocks.MockPromotionRepository)
		mockService := new(mocks.MockPromotionService)

		promotionStartDate, _ := time.Parse(time.RFC3339, "2024-03-20T00:00:00+07:00")
		promotionEndDate, _ := time.Parse(time.RFC3339, "2024-04-20T23:59:59+07:00")
		createdAt, _ := time.Parse(time.RFC3339, "2024-02-29T15:29:18.047485+07:00")
		updateAt, _ := time.Parse(time.RFC3339, "2024-02-29T15:29:18.047485+07:00")

		expectedPromo := models.Promotion{
			ID:                 1,
			PromotionID:        "cae8651b",
			PromotionName:      "Ramadhan Sale",
			DiscountType:       "Percentage",
			DiscountValue:      10.5,
			PromotionStartDate: promotionStartDate,
			PromotionEndDate:   promotionEndDate,
			CreatedAt:          createdAt,
			UpdatedAt:          updateAt,
			// DeletedAt:          nil, // Assuming soft deletion is disabled
		}

		mockRepo.On("GetPromotionbyPromotionID", "c373b046").Return(expectedPromo, nil)
		mockService.On("GetPromotionbyPromotionID", "c373b046").Return(expectedPromo, nil)

		// Create handler with mocks
		e := echo.New()
		handler := handlers.PSQLGetPromotionbyPromotionID(mockService)
		e.GET("/getpromotion/:promotion_id", handler)

		// Create request
		req := httptest.NewRequest(http.MethodGet, "/getpromotion/cae8651b", nil)
		rec := httptest.NewRecorder()

		// Execute request
		e.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualPromo models.Promotion
		err := json.Unmarshal(rec.Body.Bytes(), &actualPromo)
		assert.NoError(t, err)
		assert.Equal(t, expectedPromo, actualPromo)

		// Assert mock interactions
		mockRepo.AssertExpectations(t)
		mockService.AssertExpectations(t)
	})

	t.Run("Promotion not found", func(t *testing.T) {
		// ... (set up mocks for error case)
		mockRepo := new(mocks.MockPromotionRepository)
		mockService := new(mocks.MockPromotionService)

		mockPromoErr := errors.New("promotion not found")

		mockRepo.On("GetPromotionbyPromotionID", "cb7360g6").Return(models.Promotion{}, mockPromoErr)
		mockService.On("GetPromotionbyPromotionID", "cb7360g6").Return(models.Promotion{}, &exception.PromotionIDNotFoundError{
			Message:     "Promotion Not Found",
			PromotionID: "cb7360g6",
		})

		// Create handler with mocks
		e := echo.New()
		handler := handlers.PSQLGetPromotionbyPromotionID(mockService)
		e.GET("/getpromotion/:promotion_id", handler)

		// Create request
		req := httptest.NewRequest(http.MethodGet, "/getpromotion/cb7360g6", nil)
		rec := httptest.NewRecorder()

		// Execute request
		e.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusNotFound, rec.Code)

		// Assert mock interactions
		mockRepo.AssertExpectations(t)
		mockService.AssertExpectations(t)
	})
}

func TestPSQLUpdatePromotionbyPromotionID(t *testing.T) {
	t.Run("Successful update of promotion", func(t *testing.T) {
		// Set up mocks
		mockRepo := new(mocks.MockPromotionRepository)
		mockService := new(mocks.MockPromotionService)

		existingPromo := models.Promotion{
			PromotionID:        "cae8651b",
			PromotionName:      "Ramadhan Sale",
			DiscountType:       "percentage",
			DiscountValue:      10.5,
			PromotionStartDate: time.Now(),
			PromotionEndDate:   time.Now().Add(24 * time.Hour),
		}

		updatedPromo := existingPromo
		updatedPromo.PromotionName = "Winter Sale"
		updatedPromo.DiscountType = "fixed"
		updatedPromo.DiscountValue = 15.0

		mockRepo.On("GetPromotionbyPromotionID", "c373b046").Return(existingPromo, nil)
		mockRepo.On("UpdatePromotionbyPromotionID", updatedPromo).Return(updatedPromo, nil)
		mockService.On("UpdatePromotionbyPromotionID", updatedPromo).Return(updatedPromo, nil)

		// Create handler with mocks
		e := echo.New()
		handler := handlers.PSQLUpdatePromotionbyPromotionID(mockService)
		e.PUT("/updatepromotion/:promotion_id", handler)

		// Create request
		reqBody, err := json.Marshal(updatedPromo)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPut, "/updatepromotion/c373b046", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		// Execute request
		e.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusOK, rec.Code)
		var actualPromo models.Promotion
		err = json.Unmarshal(rec.Body.Bytes(), &actualPromo)
		assert.NoError(t, err)
		assert.Equal(t, updatedPromo, actualPromo)

		// Assert mock interactions
		mockRepo.AssertExpectations(t)
		mockService.AssertExpectations(t)
	})

	t.Run("Promotion not found", func(t *testing.T) {
		// ... (set up mocks for error case)
		mockRepo := new(mocks.MockPromotionRepository)
		mockService := new(mocks.MockPromotionService)

		// Create request
		reqBody, err := json.Marshal(models.Promotion{}) // Empty body to trigger binding error
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPut, "/updatepromotion/non_existent_id", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		handler := handlers.PSQLUpdatePromotionbyPromotionID(mockService)
		e.PUT("/updatepromotion/:promotion_id", handler)

		// Execute request
		e.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusNotFound, rec.Code) // Expected error code for non-existent promotion

		// Assert mock interactions
		mockRepo.AssertExpectations(t)
		mockService.AssertExpectations(t)
	})

	t.Run("Invalid promotion data", func(t *testing.T) {
		// ... (set up mocks for error case)
		mockRepo := new(mocks.MockPromotionRepository)
		mockService := new(mocks.MockPromotionService)

		// Create request
		req := httptest.NewRequest(http.MethodPut, "/updatepromotion/c373b046", nil) // Empty body
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		handler := handlers.PSQLUpdatePromotionbyPromotionID(mockService)
		e.PUT("/updatepromotion/:promotion_id", handler)

		// Execute request
		e.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusBadRequest, rec.Code) // Expected error code for invalid data

		// Assert mock interactions
		mockRepo.AssertExpectations(t)
		mockService.AssertExpectations(t)
	})
}

func TestPSQLDeletePromotionbyPromotionID(t *testing.T) {
	t.Run("Successful deletion of promotion", func(t *testing.T) {
		// Set up mocks
		mockRepo := new(mocks.MockPromotionRepository)
		mockService := new(mocks.MockPromotionService)

		mockRepo.On("DeletePromotionbyPromotionID", "valid_id").Return(nil)
		mockService.On("DeletePromotionbyPromotionID", "valid_id").Return(nil)

		// Create handler with mocks
		e := echo.New()
		handler := handlers.PSQLDeletePromotionbyPromotionID(mockService)
		e.DELETE("/deletepromotion/:promotion_id", handler)

		// Create request
		req := httptest.NewRequest(http.MethodDelete, "/deletepromotion/valid_id", nil)
		rec := httptest.NewRecorder()

		// Execute request
		e.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Empty(t, rec.Body.String()) // No content expected for 204

		// Assert mock interactions
		mockRepo.AssertExpectations(t)
		mockService.AssertExpectations(t)
	})

	t.Run("Promotion not found", func(t *testing.T) {
		// Set up mocks
		mockRepo := new(mocks.MockPromotionRepository)
		mockService := new(mocks.MockPromotionService)

		mockRepo.On("DeletePromotionbyPromotionID", "non_existent_id").Return(&exception.PromotionIDNotFoundError{Message: "Promotion Not Found", PromotionID: "non_existent_id"})
		mockService.On("DeletePromotionbyPromotionID", "non_existent_id").Return(&exception.PromotionIDNotFoundError{Message: "Promotion Not Found", PromotionID: "non_existent_id"})

		// Create request
		req := httptest.NewRequest(http.MethodDelete, "/promotions/non_existent_id", nil)
		rec := httptest.NewRecorder()

		e := echo.New()
		handler := handlers.PSQLDeletePromotionbyPromotionID(mockService)
		e.DELETE("/deletepromotion/:promotion_id", handler)

		// Execute request
		e.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusNotFound, rec.Code)
		// No need to assert body as there's no expected response

		// Assert mock interactions
		mockRepo.AssertExpectations(t)
		mockService.AssertExpectations(t)
	})
}
