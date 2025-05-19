package v1

import (
	"dg-server/internal/domain"
	"dg-server/internal/usecase"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type DepartmentHandler struct {
	uc *usecase.DepartmentUseCase
}

func NewDepartmentHandler(uc *usecase.DepartmentUseCase) *DepartmentHandler {
	return &DepartmentHandler{uc: uc}
}

func (h *DepartmentHandler) CreateDepartment(c echo.Context) error {
	var dept domain.Department
	if err := c.Bind(&dept); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	// If client doesn't provide UUID, generate it
	if dept.ID == uuid.Nil {
		dept.ID = uuid.New()
	}

	if dept.OrganizationID == uuid.Nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "organization_id is required"})
	}

	if dept.DepartmentName == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "department_name is required"})
	}

	if err := h.uc.CreateDepartment(c.Request().Context(), &dept); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, dept)
}
