package handler

import (
	"net/http"

	"github.com/BIC-Final-Project/backend/internal/operational/entity"
	"github.com/BIC-Final-Project/backend/internal/operational/usecase"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type MembershipTypeHandler struct {
	u usecase.MembershipTypeUsecase
}

func (h *MembershipTypeHandler) CreateMembershipType(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	var input entity.CreateMembershipType
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	data, err := h.u.InsertMembershipType(c.Context(), input)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Membership Type created",
		"data":    data,
	})
}

func (h *MembershipTypeHandler) UpdateMembershipType(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	id := c.Params("id")
	var input entity.UpdateMembershipType
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid input")
	}

	data, err := h.u.UpdateMembershipType(c.Context(), id, input)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Membership Type updated",
		"data":    data,
	})
}

func (h *MembershipTypeHandler) GetAllMembershipType(c *fiber.Ctx) error {
	jenisPaket := c.Query("jenis-paket")

	data, err := h.u.FindAllMembershipType(c.Context(), jenisPaket)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Membership Type found",
		"data":    data,
	})
}

func (h *MembershipTypeHandler) GetMembershipType(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := h.u.FindMembershipType(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Membership Type found",
		"data":    data,
	})
}

func (h *MembershipTypeHandler) DeleteMembershipType(c *fiber.Ctx) error {
	_, _, err := utils.GetCurrentAuthUser(c)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, err.Error())
	}

	id := c.Params("id")
	err = h.u.DeleteMembershipType(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Membership Type deleted",
	})
}

func NewMembershipTypeHandler(u usecase.MembershipTypeUsecase) *MembershipTypeHandler {
	return &MembershipTypeHandler{u}
}
