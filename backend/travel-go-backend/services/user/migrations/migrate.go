package migrations

import (
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	place_entity "gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/entity"
	review_entity "gitlab.com/virtual-travel/travel-go-backend/services/place/domain/review/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/entity"
)

func AutoMigrateDB(db *pg.Database) error {
	err := db.Gorm.AutoMigrate(
		entity.User{}, place_entity.Place{},
		place_entity.PlaceType{}, place_entity.PlaceImg{},
		review_entity.Review{}, review_entity.ReviewImg{},
		review_entity.ChildReview{}, review_entity.ChildReviewImg{},
	)
	if err != nil {
		return err
	}
	logger.Debug("Migrate successfully.")
	return nil
}
