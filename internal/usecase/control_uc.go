package usecase

import (
	"context"
	"dg-server/internal/domain"
	"dg-server/internal/repository"
	"errors"

	"github.com/google/uuid"
)

type ControlUseCase struct {
	crepo repository.ControlRepository
}

type ControlStatusSummary struct {
	Total         int `json:"total"`
	Compliant     int `json:"compliant"`
	NonCompliant  int `json:"non_compliant"`
	NotApplicable int `json:"not_applicable"`
}

func NewControlUseCase(crepo repository.ControlRepository) *ControlUseCase {
	return &ControlUseCase{crepo: crepo}
}

func (uc *ControlUseCase) CreateControl(ctx context.Context, c *domain.Control) error {
	if !domain.IsValidStatus(c.Status) {
		return errors.New("invalid status: must be Compliant, NonCompliant, or NotApplicable")
	}

	return uc.crepo.Create(ctx, c)
}

func (uc *ControlUseCase) GetControlById(ctx context.Context, id uuid.UUID) (*domain.Control, error) {
	return uc.crepo.FindByID(ctx, id)
}

func (uc *ControlUseCase) GetControlByCode(ctx context.Context, code string) (*domain.Control, error) {
	return uc.crepo.FindByCode(ctx, code)
}

func (uc *ControlUseCase) ListControls(ctx context.Context) ([]domain.Control, error) {
	return uc.crepo.FindAll(ctx)
}

func (uc *ControlUseCase) UpdateControl(ctx context.Context, c *domain.Control) error {
	// Optional: fetch existing and apply patch logic if needed
	return uc.crepo.Update(ctx, c)
}

func (uc *ControlUseCase) DeleteControl(ctx context.Context, id uuid.UUID) error {
	return uc.crepo.Delete(ctx, id)
}

func (uc *ControlUseCase) GetControlStatusSummary(ctx context.Context) (*ControlStatusSummary, error) {
	controls, err := uc.crepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	summary := &ControlStatusSummary{}
	summary.Total = len(controls)

	for _, c := range controls {
		switch c.Status {
		case domain.StatusCompliant:
			summary.Compliant++
		case domain.StatusNonCompliant:
			summary.NonCompliant++
		case domain.StatusNotApplicable:
			summary.NotApplicable++
		}
	}

	return summary, nil
}
