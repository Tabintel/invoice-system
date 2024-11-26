package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Tabintel/invoice-system/internal/service"
)

type DashboardHandler struct {
    dashboardService *service.DashboardService
}

// @Summary Get dashboard statistics
// @Description Get overview of invoices and payments
// @Tags dashboard
// @Security BearerAuth
// @Produce json
// @Success 200 {object} DashboardResponse
// @Router /dashboard [get]
func (h *DashboardHandler) GetStats(c *fiber.Ctx) error {
    stats, err := h.dashboardService.GetStats(c.Context())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch dashboard stats",
        })
    }
    return c.JSON(stats)
}
