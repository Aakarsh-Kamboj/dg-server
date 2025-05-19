//go:build wireinject
// +build wireinject

package main

import (
	"dg-server/config"
	"dg-server/infrastructure"
	"dg-server/infrastructure/database"
	"dg-server/internal/repository"
	v1 "dg-server/internal/transport/http/v1"
	"dg-server/internal/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProvideGormDB makes the inner *gorm.DB available for DI.
func ProvideGormDB(d *database.Database) *gorm.DB {
	return d.DB
}

// ProvideUnitOfWorkFactory gives Wire a func() repository.UnitOfWork to inject
// into your NewOrgRegistrationUseCase.
func ProvideUnitOfWorkFactory(db *gorm.DB) func() repository.UnitOfWork {
	return func() repository.UnitOfWork {
		return repository.NewUnitOfWork(db)
	}
}

// InitializeServer wires up Echo → MyService → Server.
func InitializeServer() (*Server, error) {
	wire.Build(
		// infrastructure:
		config.ProvideConfig,
		infrastructure.ProvideLogger,

		// database layer:
		database.NewDatabase,
		ProvideGormDB,

		// repositories:
		//repository.NewTenantRepository,
		//repository.NewOrganizationRepository,
		//repository.NewUserRepository,
		repository.NewControlRepository,
		repository.NewEvidenceTaskRepository,
		repository.NewFrameworkRepository,
		repository.NewDepartmentRepository,
		ProvideUnitOfWorkFactory,

		// use-case:
		usecase.NewOrgRegistrationUseCase,
		usecase.NewControlUseCase,
		usecase.NewEvidenceTaskUseCase,
		usecase.NewFrameworkUseCase,
		usecase.NewDepartmentUseCase,

		// handler:
		v1.NewOnboardingHandler,
		v1.NewControlHandler,
		v1.NewEvidenceTaskHandler,
		v1.NewFrameworkHandler,
		v1.NewDepartmentHandler,

		// framework:
		NewEcho,
		NewServer,
	)
	return &Server{}, nil
}
