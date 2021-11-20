package app

import (
	"fiber_blog/config"
	v1 "fiber_blog/internal/controller/http/v1"
	"fiber_blog/internal/usecase/repo"
	"fiber_blog/pkg/logger"
	"fiber_blog/pkg/mapStorage"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	ms := mapStorage.New()

	// Use case
	blogUseCase := repo.New(ms)

	// HTTP Server
	server := fiber.New()
	v1.NewRouter(server, l, blogUseCase)

	go server.Listen(":" + cfg.Port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	}
}
