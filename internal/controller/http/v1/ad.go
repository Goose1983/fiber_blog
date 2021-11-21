package v1

import (
	"encoding/json"
	"fiber_blog/internal/entity"
	"fiber_blog/internal/usecase"
	"fiber_blog/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

var advertiseUseCase usecase.Advertise
var advertiseLog logger.Interface

func newAdRoutes(h fiber.Router, a usecase.Advertise, l logger.Interface) {
	advertiseUseCase = a
	advertiseLog = l

	h.Get("/ads", retrieveAd)
}

type retrieveAdResponse struct {
	entity.Ad
}

func retrieveAd(c *fiber.Ctx) error {
	ad, err := advertiseUseCase.GetAd()
	if err != nil {
		return err
	}
	response, err := json.Marshal(&retrieveAdResponse{Ad: ad})
	if err != nil {
		return err
	}
	c.Send(response)
	return nil
}
