package usecase

import (
	"context"
	"dg-server/internal/domain"
	"dg-server/internal/repository"
)

type DepartmentUseCase struct {
	repo repository.DepartmentRepository
}

func NewDepartmentUseCase(repo repository.DepartmentRepository) *DepartmentUseCase {
	return &DepartmentUseCase{repo: repo}
}

func (uc *DepartmentUseCase) CreateDepartment(ctx context.Context, dept *domain.Department) error {
	return uc.repo.Create(ctx, dept)
}
