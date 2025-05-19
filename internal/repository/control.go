package repository

import (
	"context"
	"dg-server/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ControlRepository interface {
	Create(ctx context.Context, c *domain.Control) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Control, error)
	FindByCode(ctx context.Context, code string) (*domain.Control, error)
	FindAll(ctx context.Context) ([]domain.Control, error)
	Update(ctx context.Context, c *domain.Control) error
	Delete(ctx context.Context, id uuid.UUID) error
	AddControlsToFramework(ctx context.Context, frameworkID uuid.UUID, controls []domain.Control) error
	RemoveControlsFromFramework(ctx context.Context, frameworkID uuid.UUID, controls []domain.Control) error
	ClearControlsFromFramework(ctx context.Context, frameworkID uuid.UUID) error
}

type controlRepository struct {
	db *gorm.DB
}

func NewControlRepository(db *gorm.DB) ControlRepository {
	return &controlRepository{db: db}
}

func (r *controlRepository) Create(ctx context.Context, c *domain.Control) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *controlRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Control, error) {
	var control domain.Control
	if err := r.db.WithContext(ctx).Preload("Organization").First(&control, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &control, nil
}

func (r *controlRepository) FindByCode(ctx context.Context, code string) (*domain.Control, error) {
	var control domain.Control
	if err := r.db.WithContext(ctx).First(&control, "control_code = ?", code).Error; err != nil {
		return nil, err
	}
	return &control, nil
}

func (r *controlRepository) FindAll(ctx context.Context) ([]domain.Control, error) {
	var controls []domain.Control
	if err := r.db.WithContext(ctx).Find(&controls).Error; err != nil {
		return nil, err
	}
	return controls, nil
}

func (r *controlRepository) Update(ctx context.Context, c *domain.Control) error {
	existing := domain.Control{}
	if err := r.db.WithContext(ctx).First(&existing, "id=?", c.ID).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).Model(&existing).Updates(c).Error
}

func (r *controlRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Control{}, "id = ?", id).Error
}

func (r *controlRepository) AddControlsToFramework(ctx context.Context, frameworkID uuid.UUID, controls []domain.Control) error {
	var framework domain.Framework
	if err := r.db.WithContext(ctx).First(&framework, "id = ?", frameworkID).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).Model(&framework).Association("Controls").Append(controls)
}

func (r *controlRepository) RemoveControlsFromFramework(ctx context.Context, frameworkID uuid.UUID, controls []domain.Control) error {
	var framework domain.Framework
	if err := r.db.WithContext(ctx).First(&framework, "id = ?", frameworkID).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).Model(&framework).Association("Controls").Delete(controls)
}

func (r *controlRepository) ClearControlsFromFramework(ctx context.Context, frameworkID uuid.UUID) error {
	var framework domain.Framework
	if err := r.db.WithContext(ctx).First(&framework, "id = ?", frameworkID).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).Model(&framework).Association("Controls").Clear()
}
