package repository

import (
	"context"
	"dg-server/internal/domain"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	Create(ctx context.Context, dept *domain.Department) error
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db: db}
}

func (r *departmentRepository) Create(ctx context.Context, dept *domain.Department) error {
	return r.db.Create(dept).Error
}
