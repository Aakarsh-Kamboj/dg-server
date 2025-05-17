package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"dg-server/internal/domain"
)

// OrganizationRepository defines CRUD against the Organization metadata.
type OrganizationRepository interface {
	Create(ctx context.Context, o *domain.Organization) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Organization, error)
	FindByTenantID(ctx context.Context, tenantID uuid.UUID) (*domain.Organization, error)
	FindAll(ctx context.Context) ([]domain.Organization, error)
	Update(ctx context.Context, o *domain.Organization) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type organizationRepository struct {
	db *gorm.DB
}

// NewOrganizationRepository returns an OrganizationRepository backed by GORM.
func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db: db}
}

func (r *organizationRepository) Create(ctx context.Context, o *domain.Organization) error {
	return r.db.WithContext(ctx).Create(o).Error
}

func (r *organizationRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Organization, error) {
	var o domain.Organization
	if err := r.db.WithContext(ctx).First(&o, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *organizationRepository) FindByTenantID(ctx context.Context, tenantID uuid.UUID) (*domain.Organization, error) {
	var o domain.Organization
	if err := r.db.WithContext(ctx).First(&o, "tenant_id = ?", tenantID).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *organizationRepository) FindAll(ctx context.Context) ([]domain.Organization, error) {
	var orgs []domain.Organization
	if err := r.db.WithContext(ctx).Find(&orgs).Error; err != nil {
		return nil, err
	}
	return orgs, nil
}

func (r *organizationRepository) Update(ctx context.Context, o *domain.Organization) error {
	return r.db.WithContext(ctx).Save(o).Error
}

func (r *organizationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Organization{}, "id = ?", id).Error
}
