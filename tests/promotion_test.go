package tests

import (
	"errors"
	"testing"
	"time"

	schema "smkdevid/echocommercehub/internal/models/schema"
	"smkdevid/echocommercehub/internal/services/promotions"
	"smkdevid/echocommercehub/tests/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePromotion(t *testing.T) {
	t.Run("Successful Promotion Created", func(t *testing.T) {
		mockPromotionRepo := new(mocks.MockPromotionRepository)

		userService := promotions.NewPromotionService(
			&promotions.PromotionServiceImpl{
				PromotionRepo: mockPromotionRepo,
			})

		expectedPromotion := schema.Promotion{
			PromotionID:        "cae8651b",
			PromotionName:      "Ramadhan Sale",
			DiscountType:       "percentage",
			DiscountValue:      10.5,
			PromotionStartDate: time.Now(),
			PromotionEndDate:   time.Now().Add(24 * time.Hour),
		}

		mockPromotionRepo.On("CreatePromotion", mock.AnythingOfType("schema.Promotion")).Return(expectedPromotion, nil)

		results, err := userService.CreatePromotion(expectedPromotion)
		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Equal(t, expectedPromotion, results)
		mockPromotionRepo.AssertExpectations(t)
	})

	t.Run("Error on Promotion Creation", func(t *testing.T) {
		mockPromotionRepo := new(mocks.MockPromotionRepository)

		userService := promotions.NewPromotionService(
			&promotions.PromotionServiceImpl{
				PromotionRepo: mockPromotionRepo,
			})

		expectedPromotion := schema.Promotion{
			PromotionID:        "cae8651b",
			PromotionName:      "Ramadhan Sale",
			DiscountType:       "percentage",
			DiscountValue:      10.5,
			PromotionStartDate: time.Now(),
			PromotionEndDate:   time.Now().Add(24 * time.Hour),
		}

		expectedErr := errors.New("failed to create promotion")
		mockPromotionRepo.On("CreatePromotion", mock.AnythingOfType("schema.Promotion")).Return(schema.Promotion{}, expectedErr)

		results, err := userService.CreatePromotion(expectedPromotion)
		assert.Error(t, expectedErr)
		assert.NotNil(t, results)
		assert.Equal(t, expectedErr, err)
		mockPromotionRepo.AssertExpectations(t)
	})
}

func TestPSQLGetAllPromotionData(t *testing.T) {
	t.Run("Successful Retrieval of Promotions", func(t *testing.T) {
		// Set up mocks
		mockPromotionRepo := new(mocks.MockPromotionRepository)

		userService := promotions.NewPromotionService(
			&promotions.PromotionServiceImpl{
				PromotionRepo: mockPromotionRepo,
			})

		expectedPromotion := []schema.Promotion{
			{
				PromotionID:        "cae8651b",
				PromotionName:      "Ramadhan Sale",
				DiscountType:       "percentage",
				DiscountValue:      10.5,
				PromotionStartDate: time.Now(),
				PromotionEndDate:   time.Now().Add(24 * time.Hour),
			},
		}

		mockPromotionRepo.On("GetAllPromotions").Return(expectedPromotion, nil)

		results, err := userService.GetAllPromotions()
		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Equal(t, expectedPromotion, results)
		mockPromotionRepo.AssertExpectations(t)
	})

	t.Run("Error Retrieving Promotions", func(t *testing.T) {
		// Set up mocks
		mockPromotionRepo := new(mocks.MockPromotionRepository)

		userService := promotions.NewPromotionService(
			&promotions.PromotionServiceImpl{
				PromotionRepo: mockPromotionRepo,
			})

		expectedErr := errors.New("Failed to Get Promotions")
		mockPromotionRepo.On("GetAllPromotions").Return([]schema.Promotion{}, expectedErr)

		results, err := userService.GetAllPromotions()

		assert.Error(t, err)
		assert.Empty(t, results)
		assert.Equal(t, expectedErr, err)

		// Assert mock interactions
		mockPromotionRepo.AssertExpectations(t)
	})
}

func TestPSQLGetPromotionbyPromotionID(t *testing.T) {
	t.Run("Successful Retrieval of Promotion by Promotion ID", func(t *testing.T) {
		// Set up mocks
		mockPromotionRepo := new(mocks.MockPromotionRepository)

		userService := promotions.NewPromotionService(
			&promotions.PromotionServiceImpl{
				PromotionRepo: mockPromotionRepo,
			})

		expectedPromotion := schema.Promotion{
			PromotionID:        "cae8651b",
			PromotionName:      "Ramadhan Sale",
			DiscountType:       "percentage",
			DiscountValue:      10.5,
			PromotionStartDate: time.Now(),
			PromotionEndDate:   time.Now().Add(24 * time.Hour),
		}

		mockPromotionRepo.On("GetPromotionbyPromotionID", "cae8651b").Return(expectedPromotion, nil)

		results, err := userService.GetPromotionbyPromotionID("cae8651b")
		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Equal(t, expectedPromotion, results)
		mockPromotionRepo.AssertExpectations(t)
	})

	t.Run("Promotion not found", func(t *testing.T) {
		mockPromotionRepo := new(mocks.MockPromotionRepository)

		userService := promotions.NewPromotionService(
			&promotions.PromotionServiceImpl{
				PromotionRepo: mockPromotionRepo,
			})

		expectedErr := errors.New("Promotion not Found")
		mockPromotionRepo.On("GetPromotionbyPromotionID", "cb7360g6").Return(schema.Promotion{}, expectedErr)

		results, err := userService.GetPromotionbyPromotionID("cb7360g6")

		assert.Error(t, err)
		assert.Empty(t, results)
		assert.Equal(t, expectedErr, err)

		// Assert mock interactions
		mockPromotionRepo.AssertExpectations(t)
	})
}

func TestPSQLUpdatePromotionbyPromotionID(t *testing.T) {
	t.Run("Successful Update the Promotion", func(t *testing.T) {
		// Set up mocks
		mockPromotionRepo := new(mocks.MockPromotionRepository)

		userService := promotions.NewPromotionService(
			&promotions.PromotionServiceImpl{
				PromotionRepo: mockPromotionRepo,
			})

		existingPromo := schema.Promotion{
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

		mockPromotionRepo.On("UpdatePromotionbyPromotionID", updatedPromo).Return(updatedPromo, nil)

		results, err := userService.UpdatePromotionbyPromotionID(updatedPromo)
		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Equal(t, updatedPromo, results)
		mockPromotionRepo.AssertExpectations(t)
	})

	t.Run("Failed to Update the Promotion", func(t *testing.T) {
		// ... (set up mocks for error case)
		mockPromotionRepo := new(mocks.MockPromotionRepository)

		userService := promotions.NewPromotionService(
			&promotions.PromotionServiceImpl{
				PromotionRepo: mockPromotionRepo,
			})

		existingPromo := schema.Promotion{
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

		expectedErr := errors.New("Failed to Update Promotion")
		mockPromotionRepo.On("UpdatePromotionbyPromotionID", updatedPromo).Return(schema.Promotion{}, expectedErr)

		results, err := userService.UpdatePromotionbyPromotionID(updatedPromo)
		assert.Error(t, err)
		assert.Empty(t, results)
		assert.Equal(t, expectedErr, err)

		mockPromotionRepo.AssertExpectations(t)
	})
}

func TestPSQLDeletePromotionbyPromotionID(t *testing.T) {
	t.Run("Successful Deletion of Promotion", func(t *testing.T) {
		// Set up mocks
		mockPromotionRepo := new(mocks.MockPromotionRepository)

		userService := promotions.NewPromotionService(
			&promotions.PromotionServiceImpl{
				PromotionRepo: mockPromotionRepo,
			})

		mockPromotionRepo.On("DeletePromotionbyPromotionID", "cb7360g6").Return(nil)

		// Assert response
		results := userService.DeletePromotionbyPromotionID("cb7360g6")
		assert.NoError(t, results)
		mockPromotionRepo.AssertExpectations(t)
	})

	t.Run("Promotion not found", func(t *testing.T) {
		// Set up mocks
		mockPromotionRepo := new(mocks.MockPromotionRepository)

		userService := promotions.NewPromotionService(
			&promotions.PromotionServiceImpl{
				PromotionRepo: mockPromotionRepo,
			})

		expectedErr := errors.New("Promotion Not Found")
		mockPromotionRepo.On("DeletePromotionbyPromotionID", "cb7360g6").Return(expectedErr)

		// Assert response
		results := userService.DeletePromotionbyPromotionID("cb7360g6")
		assert.Error(t, results)
		mockPromotionRepo.AssertExpectations(t)
	})
}
