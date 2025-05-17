package dto

import (
	"github.com/google/uuid"
)

type OrgRegistrationRequest struct {
	// — Tenant / Organization info —
	OrganizationName  string  `json:"organization_name" validate:"required"`
	Street            string  `json:"street"`
	City              string  `json:"city"             validate:"required"`
	State             string  `json:"state,omitempty"`
	PostalCode        string  `json:"postal_code,omitempty"`
	Country           string  `json:"country"          validate:"required"`
	ContactEmail      string  `json:"contact_email"    validate:"required,email"`
	OrganizationPhone *string `json:"organization_phone,omitempty"`
	OrganizationSite  *string `json:"organization_site,omitempty"`

	// — Admin user info —
	AdminEmail     string  `json:"admin_email"     validate:"required,email"`
	AdminPassword  string  `json:"admin_password"  validate:"required,min=8"`
	AdminFirstName *string `json:"admin_first_name,omitempty"`
	AdminLastName  *string `json:"admin_last_name,omitempty"`

	// — Any extra “profile” domains —
	Industry    *string `json:"industry,omitempty"`
	CompanySize *int    `json:"company_size,omitempty"`
	// …add as many as you need…
}

type OrgRegistrationResponse struct {
	TenantID uuid.UUID `json:"tenant_id"`
	UserID   uuid.UUID `json:"admin_user_id"`
}
