package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cothromachd/game/internal/game/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) StartGame(ctx *fiber.Ctx) error {
	userID, userRole := ctx.Locals("id").(int), ctx.Locals("role").(string)
	if userRole == models.WorkerRole {
		logError("StartGame", errors.New("invalid role"))
		return ctx.Status(fiber.StatusForbidden).SendString("You can't use this request")
	}

	reqBody := ctx.Body()
	startGameReq := models.StartGameRequest{}
	err := json.Unmarshal(reqBody, &startGameReq)
	if err != nil {
		logError("StartGame", err)
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid JSON input")
	}

	success, err := h.gameService.StartGame(context.Background(), userID, startGameReq.WorkerIDs, startGameReq.OrderID)
	if err != nil {
		logError("StartGame", err)
		return err
	}

	if success {
		return ctx.Status(fiber.StatusOK).SendString("You win!")
	} else {
		return ctx.Status(fiber.StatusOK).SendString("You lose!")
	}
}
