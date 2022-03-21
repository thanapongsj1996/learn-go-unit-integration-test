package handlers_test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go-unit-integration-test/handlers"
	"go-unit-integration-test/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestPromotionHandlerCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		// Arrange
		amount := 100
		expected := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		// http://localhost:8000/calculate?amount=100
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		// Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}
	})
}
