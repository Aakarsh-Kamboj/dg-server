package v1

import (
	"dg-server/internal/domain"
	"dg-server/internal/usecase"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type FrameworkHandler struct {
	uc *usecase.FrameworkUseCase
}

func NewFrameworkHandler(uc *usecase.FrameworkUseCase) *FrameworkHandler {
	return &FrameworkHandler{uc: uc}
}

func (h *FrameworkHandler) CreateFramework(c echo.Context) error {
	var framework domain.Framework
	if err := c.Bind(&framework); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	if err := c.Validate(framework); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := h.uc.CreateFramework(c.Request().Context(), &framework); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, framework)
}

func (h *FrameworkHandler) GetFrameworkByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid UUID"})
	}

	framework, err := h.uc.GetFrameworkByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, framework)
}

func (h *FrameworkHandler) GetFrameworkByName(c echo.Context) error {
	name := c.Param("name")

	framework, err := h.uc.GetFrameworkByName(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, framework)
}

func (h *FrameworkHandler) ListFrameworks(c echo.Context) error {
	frameworks, err := h.uc.ListFrameworks(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, frameworks)
}

func (h *FrameworkHandler) UpdateFramework(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid UUID"})
	}

	var framework domain.Framework
	if err := c.Bind(&framework); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	framework.ID = id // enforce ID

	if err := c.Validate(framework); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := h.uc.UpdateFramework(c.Request().Context(), &framework); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, framework)
}

func (h *FrameworkHandler) DeleteFramework(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid UUID"})
	}

	if err := h.uc.DeleteFramework(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *FrameworkHandler) GetCompliancePercentage(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid UUID"})
	}

	percentage, err := h.uc.GetCompliancePercentage(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"compliance_percentage": percentage})
}

func (h *FrameworkHandler) AddControlToFramework(c echo.Context) error {
	frameworkIDStr := c.Param("framework_id")
	controlIDStr := c.Param("control_id")

	frameworkID, err := uuid.Parse(frameworkIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid framework UUID"})
	}

	controlID, err := uuid.Parse(controlIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid control UUID"})
	}

	if err := h.uc.AddControlToFramework(c.Request().Context(), frameworkID, controlID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (h *FrameworkHandler) RemoveControlFromFramework(c echo.Context) error {
	frameworkIDStr := c.Param("framework_id")
	controlIDStr := c.Param("control_id")

	frameworkID, err := uuid.Parse(frameworkIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid framework UUID"})
	}

	controlID, err := uuid.Parse(controlIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid control UUID"})
	}

	if err := h.uc.RemoveControlFromFramework(c.Request().Context(), frameworkID, controlID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (h *FrameworkHandler) GetEvidenceTaskPercentage(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid UUID"})
	}

	percentage, err := h.uc.GetEvidenceTaskPercentage(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"evidence_task_percentage": percentage})
}
