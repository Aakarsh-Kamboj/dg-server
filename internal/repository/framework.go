package repository

import (
	"context"
	"dg-server/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FrameworkRepository interface {
	Create(ctx context.Context, f *domain.Framework) error
	FindById(ctx context.Context, id uuid.UUID) (*domain.Framework, error)
	FindByName(ctx context.Context, name string) (*domain.Framework, error)
	FindAll(ctx context.Context) ([]domain.Framework, error)
	Update(ctx context.Context, f *domain.Framework) error
	Delete(ctx context.Context, id uuid.UUID) error
	AddControls(ctx context.Context, frameworkID uuid.UUID, controls []domain.Control) error
	UpdateFrameworkControls(ctx context.Context, frameworkID uuid.UUID, controls []domain.Control) error
	RemoveControlsFromFramework(ctx context.Context, frameworkID uuid.UUID, controls []domain.Control) error
	ClearFrameworkControls(ctx context.Context, frameworkID uuid.UUID) error
}

type frameworkRepository struct {
	db *gorm.DB
}

// Returns a framework repository
func NewFrameworkRepository(db *gorm.DB) FrameworkRepository {
	return &frameworkRepository{db: db}
}

func (r *frameworkRepository) Create(ctx context.Context, f *domain.Framework) error {
	return r.db.WithContext(ctx).Create(f).Error
}

func (r *frameworkRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.Framework, error) {
	var f domain.Framework
	if err := r.db.WithContext(ctx).Preload("Controls").First(&f, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *frameworkRepository) FindByName(ctx context.Context, name string) (*domain.Framework, error) {
	var f domain.Framework
	if err := r.db.WithContext(ctx).First(&f, "framework_name=?", name).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *frameworkRepository) FindAll(ctx context.Context) ([]domain.Framework, error) {
	var frameworks []domain.Framework
	if err := r.db.WithContext(ctx).Find(&frameworks).Error; err != nil {
		return nil, err
	}
	return frameworks, nil
}

func (r *frameworkRepository) Update(ctx context.Context, f *domain.Framework) error {
	if err := r.db.WithContext(ctx).Save(&f).Error; err != nil {
		return err
	}
	return nil
}

func (r *frameworkRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Framework{}, "id=?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *frameworkRepository) AddControls(ctx context.Context, frameworkID uuid.UUID, controls []domain.Control) error {
	var framework domain.Framework
	if err := r.db.WithContext(ctx).First(&framework, "id = ?", frameworkID).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).
		Model(&framework).
		Association("Controls").
		Append(controls)
}

func (r *frameworkRepository) UpdateFrameworkControls(ctx context.Context, frameworkID uuid.UUID, controls []domain.Control) error {
	var framework domain.Framework
	if err := r.db.WithContext(ctx).First(&framework, "id = ?", frameworkID).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).Model(&framework).Association("Controls").Replace(controls)
}

func (r *frameworkRepository) RemoveControlsFromFramework(ctx context.Context, frameworkID uuid.UUID, controls []domain.Control) error {
	var framework domain.Framework
	if err := r.db.WithContext(ctx).First(&framework, "id = ?", frameworkID).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).Model(&framework).Association("Controls").Delete(controls)
}

func (r *frameworkRepository) ClearFrameworkControls(ctx context.Context, frameworkID uuid.UUID) error {
	var framework domain.Framework
	if err := r.db.WithContext(ctx).First(&framework, "id = ?", frameworkID).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).Model(&framework).Association("Controls").Clear()
}
