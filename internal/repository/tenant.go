package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"dg-server/internal/domain"
)

// TenantRepository defines CRUD against the Tenant aggregate.
type TenantRepository interface {
	Create(ctx context.Context, t *domain.Tenant) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Tenant, error)
	FindByName(ctx context.Context, name string) (*domain.Tenant, error)
	FindAll(ctx context.Context) ([]domain.Tenant, error)
	Update(ctx context.Context, t *domain.Tenant) error
	Delete(ctx context.Context, id uuid.UUID) error
	UpdateCreatedBy(ctx context.Context, tenantID, userID uuid.UUID) error
}

type tenantRepository struct {
	db *gorm.DB
}

// NewTenantRepository returns a TenantRepository backed by GORM.
func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) Create(ctx context.Context, t *domain.Tenant) error {
	return r.db.WithContext(ctx).Create(t).Error
}

func (r *tenantRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Tenant, error) {
	var t domain.Tenant
	if err := r.db.WithContext(ctx).First(&t, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *tenantRepository) FindByName(ctx context.Context, name string) (*domain.Tenant, error) {
	var t domain.Tenant
	if err := r.db.WithContext(ctx).First(&t, "tenant_name = ?", name).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *tenantRepository) FindAll(ctx context.Context) ([]domain.Tenant, error) {
	var tenants []domain.Tenant
	if err := r.db.WithContext(ctx).Find(&tenants).Error; err != nil {
		return nil, err
	}
	return tenants, nil
}

func (r *tenantRepository) Update(ctx context.Context, t *domain.Tenant) error {
	return r.db.WithContext(ctx).Save(t).Error
}

func (r *tenantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Tenant{}, "id = ?", id).Error
}

func (r *tenantRepository) UpdateCreatedBy(ctx context.Context, tenantID, userID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&domain.Tenant{}).
		Where("id = ?", tenantID).
		Update("created_by_user_id", userID).
		Error
}
