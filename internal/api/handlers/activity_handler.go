package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Tabintel/invoice-system/internal/service"
)

type ActivityHandler struct {
    activityService *service.ActivityService
}

// @Summary Get recent activities
// @Description Get list of recent invoice activities
// @Tags activities
// @Security BearerAuth
// @Produce json
// @Success 200 {array} ent.ActivityLog
// @Router /activities [get]
func (h *ActivityHandler) GetRecentActivities(c *fiber.Ctx) error {
    activities, err := h.activityService.GetRecent(c.Context())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch activities",
        })
    }
    return c.JSON(activities)
}
