package v1

import (
	"dg-server/internal/domain"
	"dg-server/internal/usecase"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EvidenceTaskHandler struct {
	uc *usecase.EvidenceTaskUseCase
}

func NewEvidenceTaskHandler(uc *usecase.EvidenceTaskUseCase) *EvidenceTaskHandler {
	return &EvidenceTaskHandler{uc: uc}
}

func (h *EvidenceTaskHandler) CreateEvidenceTask(c echo.Context) error {
	var task domain.EvidenceTask
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	if err := c.Validate(task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := h.uc.CreateEvidenceTask(c.Request().Context(), &task); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *EvidenceTaskHandler) GetEvidenceTaskByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid UUID"})
	}

	task, err := h.uc.GetEvidenceTaskByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *EvidenceTaskHandler) ListEvidenceTasks(c echo.Context) error {
	tasks, err := h.uc.GetAllEvidenceTasks(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *EvidenceTaskHandler) UpdateEvidenceTask(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid UUID"})
	}

	var task domain.EvidenceTask
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	task.ID = id // Ensure consistency

	if err := c.Validate(task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := h.uc.UpdateEvidenceTask(c.Request().Context(), &task); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *EvidenceTaskHandler) DeleteEvidenceTask(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid UUID"})
	}

	if err := h.uc.DeleteEvidenceTask(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *EvidenceTaskHandler) GetStatusSummary(c echo.Context) error {
	summary, err := h.uc.GetEvidenceStatusSummary(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, summary)
}
