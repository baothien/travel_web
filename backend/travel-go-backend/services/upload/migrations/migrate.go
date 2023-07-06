package migrations

import (
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/domain/places/entity"
)

func AutoMigrateDB(db *pg.Database) error {
	err := db.Gorm.AutoMigrate(
		entity.File{},
	)
	if err != nil {
		return err
	}
	logger.Debug("Migrate successfully")
	return nil
}
