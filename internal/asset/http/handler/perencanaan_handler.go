package handler

import (
	"net/http"

	"github.com/BIC-Final-Project/backend/internal/asset/entity"
	"github.com/BIC-Final-Project/backend/internal/asset/usecase"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type PerencanaanHandler struct {
	PerencanaanUsecase usecase.PerencanaanUsecase
}

func (h *PerencanaanHandler) CreatePerencanaan(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	var input entity.CreatePerencanaan
	err = c.BodyParser(&input)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid input")
	}

	data, err := h.PerencanaanUsecase.InsertRencana(c.Context(), input)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Rencana created",
		"data":    data,
	})
}

func (h *PerencanaanHandler) UpdatePerencanaan(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	id := c.Params("id")
	var input entity.UpdatePerencanaan
	if err = c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid input")
	}

	data, err := h.PerencanaanUsecase.UpdateRencana(c.Context(), id, input)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Rencana updated",
		"data":    data,
	})
}

func (h *PerencanaanHandler) GetAllRencana(c *fiber.Ctx) error {
	limit := c.QueryInt("limit")
	page := c.QueryInt("page")

	data, paginationData, err := h.PerencanaanUsecase.FindAllRencana(c.Context(), limit, page)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":     http.StatusOK,
		"message":    "Rencana found",
		"pagination": paginationData,
		"data":       data,
	})
}

func (h *PerencanaanHandler) GetRencana(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := h.PerencanaanUsecase.FindRencana(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Rencana found",
		"data":    data,
	})
}

func (h *PerencanaanHandler) DeleteRencana(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.PerencanaanUsecase.DeleteRencana(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Rencana deleted",
	})
}

func NewPerencanaanHandler(PerencanaanUsecase usecase.PerencanaanUsecase) *PerencanaanHandler {
	return &PerencanaanHandler{
		PerencanaanUsecase: PerencanaanUsecase,
	}
}
