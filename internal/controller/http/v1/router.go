// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"fiber_blog/internal/usecase"
	"fiber_blog/pkg/logger"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// NewRouter -.
func NewRouter(fApp *fiber.App, l logger.Interface, a usecase.Blog) {
	// Options
	fApp.Use(fiberLogger.New(fiberLogger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
	fApp.Use(recover.New())

	// Routers
	g := fApp.Group("/v1")

	newArticleRoutes(g, a, l)
	newAdRoutes(g, a, l)
}
