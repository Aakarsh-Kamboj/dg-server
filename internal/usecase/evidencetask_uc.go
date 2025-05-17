package usecase

import (
	"context"
	"dg-server/internal/domain"
	"dg-server/internal/repository"

	"github.com/google/uuid"
)

type EvidenceTaskUseCase struct {
	erepo repository.EvidenceTaskRepository
}

type EvidenceStatusSummary struct {
	Total          int `json:"total"`
	NotUploaded    int `json:"notUploaded"`
	Draft          int `json:"draft"`
	NeedsAttention int `json:"needsAttention"`
	Uploaded       int `json:"uploaded"`
}

func NewEvidenceTaskUseCase(erepo repository.EvidenceTaskRepository) *EvidenceTaskUseCase {
	return &EvidenceTaskUseCase{erepo: erepo}
}

func (uc *EvidenceTaskUseCase) CreateEvidenceTask(ctx context.Context, task *domain.EvidenceTask) error {
	return uc.erepo.Create(ctx, task)
}

func (uc *EvidenceTaskUseCase) GetEvidenceTaskByID(ctx context.Context, id uuid.UUID) (*domain.EvidenceTask, error) {
	return uc.erepo.FindByID(ctx, id)
}

func (uc *EvidenceTaskUseCase) GetAllEvidenceTasks(ctx context.Context) ([]domain.EvidenceTask, error) {
	return uc.erepo.FindAll(ctx)
}

func (uc *EvidenceTaskUseCase) UpdateEvidenceTask(ctx context.Context, task *domain.EvidenceTask) error {
	return uc.erepo.Update(ctx, task)
}

func (uc *EvidenceTaskUseCase) DeleteEvidenceTask(ctx context.Context, id uuid.UUID) error {
	return uc.erepo.Delete(ctx, id)
}

func (uc *EvidenceTaskUseCase) GetEvidenceStatusSummary(ctx context.Context) (*EvidenceStatusSummary, error) {
	tasks, err := uc.erepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	summary := &EvidenceStatusSummary{Total: len(tasks)}

	for _, t := range tasks {
		switch t.Status {
		case domain.StatusNotUploaded:
			summary.NotUploaded++
		case domain.StatusDraft:
			summary.Draft++
		case domain.StatusNeedsAttention:
			summary.NeedsAttention++
		case domain.StatusUploaded:
			summary.Uploaded++
		}
	}

	return summary, nil
}
