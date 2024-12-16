package server

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	httpServer *fiber.App
}

func (s *Server) Run(port string, handler *fiber.App) error {
	s.httpServer = handler
	return s.httpServer.Listen(":" + port)
}

func (s *Server) Shutdown() error {
	if s.httpServer == nil {
		return nil // Сервер не был запущен, ничего не делаем
	}
	return s.httpServer.Shutdown()
}
