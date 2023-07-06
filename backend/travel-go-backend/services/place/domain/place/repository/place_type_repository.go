package repository

import (
	"context"
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/entity"
	"gorm.io/gorm"
)

type PlaceTypeRepository interface {
	CreatePlaceType(ctx context.Context, place entity.PlaceType) error
	GetPlaceType(ctx context.Context) ([]entity.PlaceType, error)
}

type PlaceTypeStore struct {
	DbExecutor *gorm.DB
}

func NewPlaceTypeStoreImpl(db pg.Database) PlaceTypeRepository {
	return PlaceTypeStore{
		DbExecutor: db.Gorm,
	}
}
func (p PlaceTypeStore) CreatePlaceType(ctx context.Context, place entity.PlaceType) error {
	return p.DbExecutor.Create(&place).Error
}

func (p PlaceTypeStore) GetPlaceType(ctx context.Context) (places []entity.PlaceType, err error) {
	err = p.DbExecutor.Find(&places).Error
	return
}
