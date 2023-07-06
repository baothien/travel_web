package repository

import (
	"context"
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/checkin/entity"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gorm.io/gorm"
)

type CheckinRepository interface {
	CreateCheckin(ctx context.Context, checkin entity.Checkin) error
	ListCheckin(ctx context.Context, req request.ListCheckinReq) (pagination api_response.Pagination, err error)
}

type CheckinStore struct {
	DbExecutor *gorm.DB
}

func NewCheckinStoreImpl(db pg.Database) CheckinRepository {
	return CheckinStore{
		DbExecutor: db.Gorm,
	}
}

func (c CheckinStore) CreateCheckin(ctx context.Context, checkin entity.Checkin) error {
	return c.DbExecutor.Create(&checkin).Error
}

func (c CheckinStore) ListCheckin(ctx context.Context, req request.ListCheckinReq) (pagination api_response.Pagination, err error) {
	var review []entity.Checkin

	pagination.Limit = req.Limit
	pagination.Page = req.Page

	if req.PlaceID != "" {

	}

	exec := c.DbExecutor.Scopes(
		func(db *gorm.DB) *gorm.DB {
			if req.PlaceID == "" {
				return db
			}
			return db.Where("place_id = ?", req.PlaceID)
		},
	)

	exec.
		Scopes(api_response.Paginate(review, &pagination, exec)).
		Preload("User").
		Preload("Place").
		Order("created_at desc").
		Where("user_id = ?", req.UserID).
		Find(&review)

	pagination.Rows = review

	return pagination, nil
}
