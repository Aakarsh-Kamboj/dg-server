package usecase

import (
	"context"
	"dg-server/internal/domain"
	"dg-server/internal/repository"

	"github.com/google/uuid"
)

type FrameworkUseCase struct {
	frepo repository.FrameworkRepository
	crepo repository.ControlRepository
}

func NewFrameworkUseCase(frepo repository.FrameworkRepository, crepo repository.ControlRepository) *FrameworkUseCase {
	return &FrameworkUseCase{frepo: frepo, crepo: crepo}
}

// Create a new framework
func (uc *FrameworkUseCase) CreateFramework(ctx context.Context, f *domain.Framework) error {
	return uc.frepo.Create(ctx, f)
}

// Get a framework by ID
func (uc *FrameworkUseCase) GetFrameworkByID(ctx context.Context, id uuid.UUID) (*domain.Framework, error) {
	return uc.frepo.FindById(ctx, id)
}

// Get a framework by Name
func (uc *FrameworkUseCase) GetFrameworkByName(ctx context.Context, name string) (*domain.Framework, error) {
	return uc.frepo.FindByName(ctx, name)
}

// Get all frameworks
func (uc *FrameworkUseCase) ListFrameworks(ctx context.Context) ([]domain.Framework, error) {
	return uc.frepo.FindAll(ctx)
}

// Update a framework
func (uc *FrameworkUseCase) UpdateFramework(ctx context.Context, f *domain.Framework) error {
	return uc.frepo.Update(ctx, f)
}

// Delete a framework by ID
func (uc *FrameworkUseCase) DeleteFramework(ctx context.Context, id uuid.UUID) error {
	return uc.frepo.Delete(ctx, id)
}

// Get compliance percentage
func (uc *FrameworkUseCase) GetCompliancePercentage(ctx context.Context, frameworkID uuid.UUID) (float64, error) {
	controls, err := uc.crepo.FindByFrameworkID(ctx, frameworkID)
	if err != nil {
		return 0, err
	}
	if len(controls) == 0 {
		return 0, nil
	}

	var compliantCount int
	for _, c := range controls {
		if c.Status == domain.StatusCompliant {
			compliantCount++
		}
	}

	percentage := (float64(compliantCount) / float64(len(controls))) * 100
	return percentage, nil
}
