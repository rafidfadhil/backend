package handler

import (
	"net/http"

	"github.com/BIC-Final-Project/backend/internal/operational/entity"
	"github.com/BIC-Final-Project/backend/internal/operational/usecase"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type FasilitasHandler struct {
	u usecase.FasilitasUsecase
}

func (h *FasilitasHandler) CreateFasilitas(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	var input entity.CreateFasilitas
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid input")
	}

	input.Gambar, _ = c.FormFile("gambar")

	data, err := h.u.InsertFasilitas(c.Context(), input)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Fasilitas created",
		"data":    data,
	})
}

func (h *FasilitasHandler) UpdateFasilitas(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	id := c.Params("id")
	var input entity.UpdateFasilitas
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid input")
	}

	input.Gambar, _ = c.FormFile("gambar")

	data, err := h.u.UpdateFasilitas(c.Context(), id, input)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Fasilitas updated",
		"data":    data,
	})
}

func (h *FasilitasHandler) GetAllFasilitas(c *fiber.Ctx) error {
	limit := c.QueryInt("limit")
	page := c.QueryInt("page")

	data, paginationData, err := h.u.FindAllFasilitas(c.Context(), limit, page)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":     http.StatusOK,
		"message":    "Fasilitas found",
		"pagination": paginationData,
		"data":       data,
	})
}

func (h *FasilitasHandler) GetFasilitas(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := h.u.FindFasilitas(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Fasilitas found",
		"data":    data,
	})
}

func (h *FasilitasHandler) DeleteFasilitas(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	id := c.Params("id")
	err = h.u.DeleteFasilitas(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Fasilitas deleted",
	})
}

func (h *FasilitasHandler) GetAllFasilitasName(c *fiber.Ctx) error {
	data, err := h.u.FindAllFasilitasName(c.Context())
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Fasilitas name found",
		"data":    data,
	})
}

func NewFasilitasHandler(u usecase.FasilitasUsecase) *FasilitasHandler {
	return &FasilitasHandler{u}
}
