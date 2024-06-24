package handler

import (
	"net/http"

	"github.com/BIC-Final-Project/backend/internal/auth/entity"
	"github.com/BIC-Final-Project/backend/internal/auth/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	a usecase.AuthUsecase
}

func NewAuthHandler(a usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{a}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var input entity.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid input",
		})
	}

	ctx := c.Context()
	data, err := h.a.CreateUser(ctx, input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "User created",
		"data":    data,
	})

	return nil
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input entity.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid input",
		})
	}

	ctx := c.Context()
	userData, dataToken, err := h.a.Login(ctx, input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Login success",
		"data":    userData,
		"token":   dataToken.Token,
	})

	return nil
}
