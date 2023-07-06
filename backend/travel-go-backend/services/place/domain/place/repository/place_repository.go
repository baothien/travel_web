package repository

import (
	"context"
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/entity"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PlaceRepository interface {
	CreatePlace(ctx context.Context, place entity.Place) error
	UpdatePlace(ctx context.Context, place entity.Place) error
	GetPlace(ctx context.Context) ([]entity.Place, error)
	GetPlacePaging(ctx context.Context, param request.ListPlaceReq) (pagination api_response.Pagination, err error)
	GetPlaceById(ctx context.Context, id string) (entity.Place, error)
	FavoritePlace(ctx context.Context, param request.FavoritePlaceAddReq) error
	GetFavoritePlacePaging(ctx context.Context, param request.ListFavoritePlaceReq) (pagination api_response.Pagination, err error)
}

type PlaceStore struct {
	DbExecutor *gorm.DB
}

func NewPlaceStoreImpl(db pg.Database) PlaceRepository {
	return PlaceStore{
		DbExecutor: db.Gorm,
	}
}
func (p PlaceStore) CreatePlace(ctx context.Context, place entity.Place) error {
	err := p.DbExecutor.Create(&place).Error
	return err
}

func (p PlaceStore) UpdatePlace(ctx context.Context, place entity.Place) error {
	err := p.DbExecutor.Updates(&place).Error
	if err != nil {
		return err
	}

	errRepl := p.DbExecutor.Model(&place).Association("PlaceImg").Replace(&place.PlaceImg)
	if errRepl != nil {
		return errRepl
	}

	errReplCheckin := p.DbExecutor.Model(&place).Association("PlaceImgCheckin").Replace(&place.PlaceImgCheckin)
	if errReplCheckin != nil {
		return errReplCheckin
	}

	return nil
}

func (p PlaceStore) FavoritePlace(ctx context.Context, param request.FavoritePlaceAddReq) error {
	if !param.IsFavorite {
		err := p.DbExecutor.
			Table("userplace_favorites").
			Where("place_id = ?", param.PlaceID).
			Delete(
				map[string]interface{}{
					"place_id": param.PlaceID,
					"user_id":  param.UserID},
			).Error
		if err != nil {
			return err
		}
	} else {
		err := p.DbExecutor.
			Table("userplace_favorites").
			Clauses(clause.OnConflict{DoNothing: true}).
			Create(
				map[string]interface{}{
					"place_id": param.PlaceID,
					"user_id":  param.UserID},
			).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (p PlaceStore) GetPlace(ctx context.Context) (places []entity.Place, err error) {
	err = p.DbExecutor.
		Preload("PlaceType").
		Preload("PlaceImg").
		Preload("Review").
		Find(&places).
		Error
	return
}

func (p PlaceStore) GetPlaceById(ctx context.Context, id string) (places entity.Place, err error) {
	err = p.DbExecutor.
		Where("places.id = ?", id).
		Joins("PlaceType").
		Preload("PlaceImg").
		Preload("PlaceImgCheckin").
		Preload("Review").
		First(&places).Error
	return
}

func (p PlaceStore) GetPlacePaging(ctx context.Context, param request.ListPlaceReq) (pagination api_response.Pagination, err error) {
	pagination.Limit = param.Limit
	pagination.Page = param.Page
	var place []*entity.Place

	exec := p.DbExecutor.Scopes()

	exec.
		Scopes(api_response.Paginate(place, &pagination, exec)).
		Preload("PlaceType").
		Preload("PlaceImg").
		Preload("PlaceImgCheckin").
		Find(&place)

	pagination.Rows = place
	return
}

func (p PlaceStore) GetFavoritePlacePaging(ctx context.Context, param request.ListFavoritePlaceReq) (pagination api_response.Pagination, err error) {
	pagination.Limit = param.Limit
	pagination.Page = param.Page
	var place []*entity.Place

	placeIDs := []string{}

	err = p.DbExecutor.Debug().
		Table("userplace_favorites").
		Select("place_id").
		Where("user_id = ?", param.UserID).
		Scan(&placeIDs).Error
	if err != nil {
	}

	exec := p.DbExecutor.Scopes()

	exec.Debug().
		Scopes(api_response.Paginate(place, &pagination, exec)).
		Preload("PlaceType").
		Preload("PlaceImg").
		Preload("PlaceImgCheckin").
		Where("id IN (?)", placeIDs).
		Find(&place)

	pagination.Rows = place
	return
}
