package database

import (
	"dg-server/internal/domain"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// dbMigrate runs database migrations for all your domain entities.
// It returns any error so that callers can handle retries or clean shutdown.
//func dbMigrate(db *gorm.DB, logger *zap.Logger) error {
//	// AutoMigrate will create tables, missing foreign keys, constraints, indexes, etc.
//	// Order here is not strictly required (GORM figures out dependencies), but
//	// listing parents first can be more readable.
//	models := []interface{}{
//		&domain.Tenant{},       // create tenants first
//		&domain.Organization{}, // then orgs (FK→tenants)
//
//		&domain.User{},                   // then users (FK→tenants)
//		&domain.UserProfile{},            // then profiles (FK→users)
//		&domain.UserPreference{},         // then prefs (FK→users)
//		&domain.UserSoftDeleteMetadata{}, // finally soft-delete metadata
//	}
//
//	if err := db.AutoMigrate(models...); err != nil {
//		logger.Error("❌ Migration failed", zap.Error(err))
//		return err
//	}
//
//	logger.Info("✅ Database migration completed successfully")
//	return nil
//}

func dbMigrate(db *gorm.DB, logger *zap.Logger) error {
	// 1) Manually ensure tenants table exists first
	if !db.Migrator().HasTable(&domain.Tenant{}) {
		if err := db.Migrator().CreateTable(&domain.Tenant{}); err != nil {
			logger.Error("failed to create tenants table", zap.Error(err))
			return err
		}
	}

	// 2) Then ensure organizations
	if !db.Migrator().HasTable(&domain.Organization{}) {
		if err := db.Migrator().CreateTable(&domain.Organization{}); err != nil {
			logger.Error("failed to create organizations table", zap.Error(err))
			return err
		}
	}

	// 3) Now auto-migrate the rest in bulk
	//    This will pick up Users (with FK→tenants), Profiles, Preferences, SoftDelete, etc.
	others := []interface{}{
		&domain.User{},
		&domain.UserProfile{},
		&domain.UserPreference{},
		&domain.UserSoftDeleteMetadata{},
		&domain.Control{},
		&domain.Department{},
		&domain.EvidenceTask{},
		&domain.Framework{},
	}
	if err := db.AutoMigrate(others...); err != nil {
		logger.Error("❌ Migration failed", zap.Error(err))
		return err
	}

	logger.Info("✅ Database migration completed successfully")
	return nil
}
