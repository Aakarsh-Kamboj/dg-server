package repository

import (
	"context"
	"dg-server/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EvidenceTaskRepository interface {
	Create(ctx context.Context, task *domain.EvidenceTask) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.EvidenceTask, error)
	FindAll(ctx context.Context) ([]domain.EvidenceTask, error)
	Update(ctx context.Context, task *domain.EvidenceTask) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type evidenceTaskRepository struct {
	db *gorm.DB
}

func NewEvidenceTaskRepository(db *gorm.DB) EvidenceTaskRepository {
	return &evidenceTaskRepository{db: db}
}

func (r *evidenceTaskRepository) Create(ctx context.Context, task *domain.EvidenceTask) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *evidenceTaskRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.EvidenceTask, error) {
	var task domain.EvidenceTask
	err := r.db.WithContext(ctx).First(&task, "id = ?", id).Error
	return &task, err
}

func (r *evidenceTaskRepository) FindAll(ctx context.Context) ([]domain.EvidenceTask, error) {
	var tasks []domain.EvidenceTask
	err := r.db.WithContext(ctx).Select("id", "status").Find(&tasks).Error
	return tasks, err
}

func (r *evidenceTaskRepository) Update(ctx context.Context, task *domain.EvidenceTask) error {
	return r.db.WithContext(ctx).Save(task).Error
}

func (r *evidenceTaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.EvidenceTask{}, "id = ?", id).Error
}
