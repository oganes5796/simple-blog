package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/oganes5796/simple-blog/pkg/logger"
	"github.com/oganes5796/simple-blog/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *fiber.App {
	app := fiber.New()

	loggers := logger.NewLogger()

	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		loggers.Info("Request log",
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"latency", time.Since(start),
		)
		return err
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Home Page!")
	})

	auth := app.Group("/auth")
	auth.Post("/sign-up", h.signUp)
	auth.Post("/sign-in", h.signIn)

	api := app.Group("/api", h.authMiddleware)
	api.Post("/articles", h.CreateArticle)
	api.Get("/articles", h.GetArticles)

	return app
}
