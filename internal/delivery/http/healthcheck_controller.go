package http

import "github.com/gofiber/fiber/v2"

type HealthCheckController struct{}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

// HealthCheck performs a basic health check of the API
// @Summary Health Check
// @Description Check if the API is running
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Router /api/healthcheck [get]
// @security none
func (hc *HealthCheckController) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Health check successful",
			"success": true,
		},
	)
}
