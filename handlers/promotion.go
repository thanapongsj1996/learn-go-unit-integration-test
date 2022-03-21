package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-unit-integration-test/services"
	"strconv"
)

type PromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

type promotionHandler struct {
	promoService services.PromotionService
}

func NewPromotionHandler(promoService services.PromotionService) PromotionHandler {
	return promotionHandler{promoService: promoService}
}

func (h promotionHandler) CalculateDiscount(c *fiber.Ctx) error {
	// http://localhost:8000/calculate?amount=100

	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	discount, err := h.promoService.CalculateDiscount(amount)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString(strconv.Itoa(discount))
}
