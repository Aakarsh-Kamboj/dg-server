package repository

import (
	"gorm.io/gorm"
)

type gormUnitOfWork struct {
	tx *gorm.DB

	tenantRepo       TenantRepository
	organizationRepo OrganizationRepository
	userRepo         UserRepository
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	tx := db.Begin()
	return &gormUnitOfWork{
		tx:               tx,
		tenantRepo:       NewTenantRepository(tx),
		organizationRepo: NewOrganizationRepository(tx),
		userRepo:         NewUserRepository(tx),
	}
}

func (u *gormUnitOfWork) Tenant() TenantRepository {
	return u.tenantRepo
}

func (u *gormUnitOfWork) Organization() OrganizationRepository {
	return u.organizationRepo
}

func (u *gormUnitOfWork) User() UserRepository {
	return u.userRepo
}

func (u *gormUnitOfWork) Commit() error {
	return u.tx.Commit().Error
}

func (u *gormUnitOfWork) Rollback() error {
	return u.tx.Rollback().Error
}
