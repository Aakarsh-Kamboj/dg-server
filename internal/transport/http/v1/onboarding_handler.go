// internal/transport/http/v2/org_registration_handler.go
package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"dg-server/internal/dto"
	"dg-server/internal/usecase"
)

type OnboardingHandler struct {
	uc *usecase.OrgRegistrationUseCase
}

// NewOnboardingHandler constructs the handler; it does NOT register routes.
func NewOnboardingHandler(uc *usecase.OrgRegistrationUseCase) *OnboardingHandler {
	return &OnboardingHandler{uc: uc}
}

// OrgRegistration binds, validates, executes the use-case, and returns JSON.
func (h *OnboardingHandler) OrgRegistration(c echo.Context) error {
	var req dto.OrgRegistrationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request payload"})
	}

	// Validate the request
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	res, err := h.uc.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, res)
}
