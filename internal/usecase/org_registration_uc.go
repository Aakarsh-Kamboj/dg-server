package usecase

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"dg-server/internal/domain"
	"dg-server/internal/dto"
	"dg-server/internal/repository"
)

// OrgRegistrationUseCase handles a one‐shot org + admin signup.
type OrgRegistrationUseCase struct {
	// factory that gives a new UoW (transaction) each Execute
	newUnitOfWork func() repository.UnitOfWork
}

// NewOrgRegistrationUseCase wires in the UoW factory.
func NewOrgRegistrationUseCase(
	uowFactory func() repository.UnitOfWork,
) *OrgRegistrationUseCase {
	return &OrgRegistrationUseCase{newUnitOfWork: uowFactory}
}

// Execute runs the full orchestration in one transaction.
func (uc *OrgRegistrationUseCase) Execute(
	ctx context.Context,
	req dto.OrgRegistrationRequest,
) (*dto.OrgRegistrationResponse, error) {
	uow := uc.newUnitOfWork()
	defer uow.Rollback() // safe: if Commit() succeeds, this is a no‐op

	// 1️⃣ Create Tenant
	tenant := &domain.Tenant{
		TenantName: req.OrganizationName,
		TenantType: domain.TenantTypeOrganization,
		Address: domain.TenantAddress{
			Street:     req.Street,
			City:       req.City,
			State:      req.State,
			PostalCode: req.PostalCode,
			Country:    req.Country,
		},
		Phone:   req.OrganizationPhone,
		Website: req.OrganizationSite,
		Email:   &req.ContactEmail,
		Status:  domain.TenantStatusActive,
	}
	if err := uow.Tenant().Create(ctx, tenant); err != nil {
		return nil, fmt.Errorf("creating tenant: %w", err)
	}

	// 2️⃣ Create Organization metadata
	org := &domain.Organization{
		TenantID:     tenant.ID,
		Name:         req.OrganizationName,
		Address:      fmt.Sprintf("%s, %s, %s, %s, %s", req.Street, req.City, req.State, req.PostalCode, req.Country),
		ContactEmail: req.ContactEmail,
	}
	if err := uow.Organization().Create(ctx, org); err != nil {
		return nil, fmt.Errorf("creating organization: %w", err)
	}

	// 3️⃣ (Optional) Persist extra profile fields
	//    If you later add a domain.OrganizationProfile and
	//    OrganizationProfileRepository, you could do:
	//
	//    prof := &domain.OrganizationProfile{
	//        TenantID:    tenant.ID,
	//        Industry:    req.Industry,
	//        CompanySize: req.CompanySize,
	//    }
	//    if err := uow.OrganizationProfile().Create(ctx, prof); err != nil {
	//        return nil, fmt.Errorf("creating org profile: %w", err)
	//    }

	// 4️⃣ Hash password & create the admin user
	pwHash, err := bcrypt.GenerateFromPassword([]byte(req.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hashing password: %w", err)
	}
	user := &domain.User{
		TenantID:     tenant.ID,
		Email:        req.AdminEmail,
		PasswordHash: string(pwHash),
		Profile: &domain.UserProfile{
			FirstName: req.AdminFirstName,
			LastName:  req.AdminLastName,
		},
		Active: true,
	}
	if err := uow.User().Create(ctx, user); err != nil {
		return nil, fmt.Errorf("creating admin user: %w", err)
	}

	// 5️⃣ Mark the tenant’s CreatedByUserID
	if err := uow.Tenant().UpdateCreatedBy(ctx, tenant.ID, user.ID); err != nil {
		return nil, fmt.Errorf("updating tenant creator: %w", err)
	}

	// 6️⃣ Commit the transaction
	if err := uow.Commit(); err != nil {
		return nil, fmt.Errorf("committing registration: %w", err)
	}

	return &dto.OrgRegistrationResponse{
		TenantID: tenant.ID,
		UserID:   user.ID,
	}, nil
}
