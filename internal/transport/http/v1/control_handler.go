package v1

import (
	"dg-server/internal/domain"
	"dg-server/internal/usecase"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ControlHandler struct {
	uc *usecase.ControlUseCase
}

func NewControlHandler(uc *usecase.ControlUseCase) *ControlHandler {
	return &ControlHandler{uc: uc}
}

func (h *ControlHandler) CreateControl(c echo.Context) error {
	var control domain.Control
	if err := c.Bind(&control); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	// Validate status (should be Compliant, NonCompliant, or NotApplicable)
	if !domain.IsValidStatus(control.Status) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid status: must be Compliant, NonCompliant, or NotApplicable"})
	}

	if err := h.uc.CreateControl(c.Request().Context(), &control); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, control)
}

func (h *ControlHandler) GetControlById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid uuid"})
	}

	control, err := h.uc.GetControlById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, control)
}

func (h *ControlHandler) GetControlByCode(c echo.Context) error {
	code := c.Param("control_code")
	control, err := h.uc.GetControlByCode(c.Request().Context(), code)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, control)
}

func (h *ControlHandler) ListControls(c echo.Context) error {
	controls, err := h.uc.ListControls(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, controls)
}

func (h *ControlHandler) UpdateControl(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid UUID"})
	}

	var control domain.Control
	if err := c.Bind(&control); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	control.ID = id // force override to prevent mismatched IDs

	if err := c.Validate(control); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := h.uc.UpdateControl(c.Request().Context(), &control); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, control)
}

func (h *ControlHandler) DeleteControl(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid UUID"})
	}

	if err := h.uc.DeleteControl(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *ControlHandler) GetStatusSummary(c echo.Context) error {
	summary, err := h.uc.GetControlStatusSummary(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, summary)
}
