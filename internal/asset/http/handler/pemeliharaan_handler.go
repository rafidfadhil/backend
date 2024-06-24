package handler

import (
	"net/http"

	"github.com/BIC-Final-Project/backend/internal/asset/entity"
	"github.com/BIC-Final-Project/backend/internal/asset/usecase"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type PemeliharaanHandler struct {
	pemeliharaanUsecase usecase.PemeliharaanUsecase
}

func (h *PemeliharaanHandler) CreatePelihara(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	var input entity.CreatePelihara
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid Input")
	}

	data, err := h.pemeliharaanUsecase.InsertPelihara(c.Context(), input)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Pemeliharaan created",
		"data":    data,
	})
}

func (h *PemeliharaanHandler) UpdatePelihara(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	id := c.Params("id")
	var input entity.UpdatePelihara
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid Input")
	}

	data, err := h.pemeliharaanUsecase.UpdatePelihara(c.Context(), id, input)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Pemeliharaan updated",
		"data":    data,
	})
}

func (h *PemeliharaanHandler) GetAllPelihara(c *fiber.Ctx) error {
	limit := c.QueryInt("limit")
	page := c.QueryInt("page")

	data, paginationData, err := h.pemeliharaanUsecase.FindAllPelihara(c.Context(), limit, page)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":     http.StatusOK,
		"message":    "Pemeliharaan found",
		"pagination": paginationData,
		"data":       data,
	})
}

func (h *PemeliharaanHandler) GetPelihara(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := h.pemeliharaanUsecase.FindPelihara(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Pemeliharaan found",
		"data":    data,
	})
}

func (h *PemeliharaanHandler) DeletePelihara(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.pemeliharaanUsecase.DeletePelihara(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Pemeliharaan deleted",
	})
}

func NewPemeliharaanHandler(usecase usecase.PemeliharaanUsecase) *PemeliharaanHandler {
	return &PemeliharaanHandler{
		pemeliharaanUsecase: usecase,
	}
}
