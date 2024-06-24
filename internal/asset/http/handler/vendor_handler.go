package handler

import (
	"net/http"

	"github.com/BIC-Final-Project/backend/internal/asset/entity"
	"github.com/BIC-Final-Project/backend/internal/asset/usecase"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type VendorHandler struct {
	vendorUsecase usecase.VendorUsecase
}

func (h *VendorHandler) CreateVendor(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	var input entity.CreateVendor
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid input")
	}

	data, err := h.vendorUsecase.InsertVendor(c.Context(), input)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Vendor created",
		"data":    data,
	})
}

func (h *VendorHandler) UpdateVendor(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	id := c.Params("vendor_id")
	var input entity.UpdateVendor
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid input")
	}

	data, err := h.vendorUsecase.UpdateVendor(c.Context(), id, input)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Vendor updated",
		"data":    data,
	})
}

func (h *VendorHandler) GetAllVendor(c *fiber.Ctx) error {
	limit := c.QueryInt("limit")
	page := c.QueryInt("page")

	data, paginationData, err := h.vendorUsecase.FindAllVendors(c.Context(), limit, page)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":     http.StatusOK,
		"message":    "Vendors found",
		"pagination": paginationData,
		"data":       data,
	})
}

func (h *VendorHandler) GetVendor(c *fiber.Ctx) error {
	id := c.Params("vendor_id")
	data, err := h.vendorUsecase.FindVendor(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Vendor found",
		"data":    data,
	})
}

func (h *VendorHandler) DeleteVendor(c *fiber.Ctx) error {
	id := c.Params("vendor_id")
	err := h.vendorUsecase.DeleteVendor(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Vendor deleted",
	})
}

func NewVendorHandler(vendorUsecase usecase.VendorUsecase) *VendorHandler {
	return &VendorHandler{
		vendorUsecase: vendorUsecase,
	}
}
