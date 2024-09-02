package score

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB, redisClient *redis.Client, redisCtx context.Context) {
	redisRepo := NewRedisRepository(redisClient, redisCtx)
	repo := NewRepository(db)
	svc := NewService(repo, redisRepo)

	handler := NewHandler(svc)

	app.Post("/score", handler.AddScore)
	app.Get("/leaderboard", handler.GetLeaderboard)
	// app.Delete("/leaderboard", handler.ResetLeaderboard)
}

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// AddScore adds a new score to the leaderboard
func (h *Handler) AddScore(c *fiber.Ctx) error {
	userIDStr := c.Query("userID")
	game := c.Query("game")
	scoreStr := c.Query("score")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid userID")
	}
	score, err := strconv.ParseFloat(scoreStr, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid score")
	}

	if err := h.Service.AddScore(uint(userID), game, score); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to submit score")
	}

	return c.SendStatus(fiber.StatusOK)
}

// GetLeaderboard retrieves the leaderboard
func (h *Handler) GetLeaderboard(c *fiber.Ctx) error {
	scores, err := h.Service.GetLeaderboard()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve leaderboard")
	}
	return c.JSON(scores)
}

// GetUserRank retrieves the rank of a user
// func (h *Handler) GetUserRank(c *fiber.Ctx) error {
// 	userIDStr := c.Query("userID")
// 	userID, err := strconv.Atoi(userIDStr)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString("Invalid userID")
// 	}

// 	rank, err := h.Service.GetUserRank(uint(userID))
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve rank")
// 	}

// 	return c.JSON(fiber.Map{"rank": rank})
// }

// ResetLeaderboard clears the leaderboard
// func (h *Handler) ResetLeaderboard(c *fiber.Ctx) error {
// 	if err := h.Service.ClearScores(); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString("Failed to reset leaderboard")
// 	}
// 	return c.SendStatus(fiber.StatusOK)
// }
