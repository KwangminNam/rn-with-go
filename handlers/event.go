package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kwangminnam/rn-with-go/models"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	events, err := h.repository.GetMany(context)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   events,
	})
}

func (h *EventHandler) GetOne(ctx *fiber.Ctx) error {
	return nil
}

func (h *EventHandler) CreateOne(ctx *fiber.Ctx) error {
	return nil
}

func NewEventHandler(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}
	router.Get("/", handler.GetMany)
	router.Get("/:eventId", handler.GetOne)
	router.Post("/", handler.CreateOne)
}

func (h *EventHandler) getMany(c *fiber.Ctx) error {
	return nil
}
