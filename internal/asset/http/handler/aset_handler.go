package handler

import (
	"net/http"

	"github.com/BIC-Final-Project/backend/internal/asset/entity"
	"github.com/BIC-Final-Project/backend/internal/asset/usecase"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type AsetHandler struct {
	asetUsecase usecase.AsetUsecase
}

func (h *AsetHandler) CreateAset(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	var input entity.CreateAset
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid input")
	}

	input.Gambar, _ = c.FormFile("gambar")

	data, err := h.asetUsecase.InsertAset(c.Context(), input)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Aset created",
		"data":    data,
	})
}

func (h *AsetHandler) UpdateAset(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	id := c.Params("aset_id")
	var input entity.UpdateAset
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid input")
	}

	input.Gambar, _ = c.FormFile("gambar")

	data, err := h.asetUsecase.UpdateAset(c.Context(), id, input)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Aset updated",
		"data":    data,
	})
}

func (h *AsetHandler) GetAllAset(c *fiber.Ctx) error {
	limit := c.QueryInt("limit")
	page := c.QueryInt("page")

	data, paginationData, err := h.asetUsecase.FindAllAset(c.Context(), limit, page)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":     http.StatusOK,
		"message":    "Aset found",
		"pagination": paginationData,
		"data":       data,
	})
}

func (h *AsetHandler) GetAset(c *fiber.Ctx) error {
	id := c.Params("aset_id")
	data, err := h.asetUsecase.FindAset(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Aset found",
		"data":    data,
	})
}

func (h *AsetHandler) DeleteAset(c *fiber.Ctx) error {
	id := c.Params("aset_id")
	err := h.asetUsecase.DeleteAset(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Aset deleted",
	})
}

func NewAsetHandler(asetUsecase usecase.AsetUsecase) *AsetHandler {
	return &AsetHandler{
		asetUsecase: asetUsecase,
	}
}
