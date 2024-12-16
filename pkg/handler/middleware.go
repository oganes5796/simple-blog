package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) authMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing authorization token"})
	}

	tokenStr := authHeader[len("Bearer "):]

	userId, err := h.services.ParseJWT(tokenStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid or expaire token"})
	}

	c.Locals("userId", userId)

	return c.Next()

}
