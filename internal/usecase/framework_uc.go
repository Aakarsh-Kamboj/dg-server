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
	erepo repository.EvidenceTaskRepository
}

func NewFrameworkUseCase(
	frepo repository.FrameworkRepository,
	crepo repository.ControlRepository,
	erepo repository.EvidenceTaskRepository,
) *FrameworkUseCase {
	return &FrameworkUseCase{frepo: frepo, crepo: crepo, erepo: erepo}
}

// Create a new framework
func (uc *FrameworkUseCase) CreateFramework(ctx context.Context, f *domain.Framework) error {
	return uc.frepo.Create(ctx, f)
}

// Get a framework by ID
func (uc *FrameworkUseCase) GetFrameworkByID(ctx context.Context, id uuid.UUID) (*domain.Framework, error) {
	framework, err := uc.frepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Calculate and attach evidence task percentage
	percentage, err := uc.GetEvidenceTaskPercentage(ctx, framework.ID)
	if err != nil {
		return nil, err
	}
	framework.EvidenceTaskPercentage = percentage
	return uc.frepo.FindById(ctx, id)
}

// Get a framework by Name
func (uc *FrameworkUseCase) GetFrameworkByName(ctx context.Context, name string) (*domain.Framework, error) {
	return uc.frepo.FindByName(ctx, name)
}

// Get all frameworks
func (uc *FrameworkUseCase) ListFrameworks(ctx context.Context) ([]domain.Framework, error) {
	frameworks, err := uc.frepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for i := range frameworks {
		percentage, err := uc.GetEvidenceTaskPercentage(ctx, frameworks[i].ID)
		if err != nil {
			return nil, err
		}
		frameworks[i].EvidenceTaskPercentage = percentage
	}

	return frameworks, nil
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
	// Fetch the framework with its controls
	framework, err := uc.frepo.FindById(ctx, frameworkID)
	if err != nil {
		return 0, err
	}

	// If no controls are associated with the framework
	if len(framework.Controls) == 0 {
		return 0, nil
	}

	// Count compliant controls
	var compliantCount int
	for _, control := range framework.Controls {
		if control.Status == domain.StatusCompliant {
			compliantCount++
		}
	}

	// Calculate the compliance percentage
	percentage := (float64(compliantCount) / float64(len(framework.Controls))) * 100
	return percentage, nil
}

// AddControlToFramework adds a control to a framework
func (uc *FrameworkUseCase) AddControlToFramework(ctx context.Context, frameworkID uuid.UUID, controlID uuid.UUID) error {
	// Retrieve the control
	control, err := uc.crepo.FindByID(ctx, controlID)
	if err != nil {
		return err
	}

	// Delegate to the repository to add the control to the framework
	return uc.frepo.AddControls(ctx, frameworkID, []domain.Control{*control})
}

// RemoveControlFromFramework removes a control from a framework
func (uc *FrameworkUseCase) RemoveControlFromFramework(ctx context.Context, frameworkID uuid.UUID, controlID uuid.UUID) error {
	// Retrieve the control
	control, err := uc.crepo.FindByID(ctx, controlID)
	if err != nil {
		return err
	}

	// Delegate to the repository to remove the control from the framework
	return uc.frepo.RemoveControlsFromFramework(ctx, frameworkID, []domain.Control{*control})
}

func (uc *FrameworkUseCase) GetEvidenceTaskPercentage(ctx context.Context, frameworkID uuid.UUID) (float64, error) {
	total, uploaded, err := uc.erepo.GetEvidenceStatsByFramework(ctx, frameworkID)
	if err != nil {
		return 0, err
	}
	if total == 0 {
		return 0, nil
	}
	return (float64(uploaded) / float64(total)) * 100, nil
}
