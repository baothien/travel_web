package repository

import (
	"context"
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/domain/notify/entity"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gorm.io/gorm"
)

type NotifyRepository interface {
	CreateNotify(ctx context.Context, place entity.Notify) error
	GetNotifyPaging(ctx context.Context, param request.ListNotifyReq) (pagination api_response.Pagination, err error)
	UpdateIsRead(ctx context.Context, req request.UpdateIsRead) error
	CountNotifyRead(ctx context.Context, userID string) (int64, error)
}

type NotifyStore struct {
	DbExecutor *gorm.DB
}

func NewNotifyStoreImpl(db pg.Database) NotifyRepository {
	return NotifyStore{
		DbExecutor: db.Gorm,
	}
}
func (p NotifyStore) CreateNotify(ctx context.Context, place entity.Notify) error {
	return p.DbExecutor.Create(&place).Error
}

func (p NotifyStore) UpdateIsRead(ctx context.Context, req request.UpdateIsRead) error {
	return p.DbExecutor.Model(&entity.Notify{}).Where("id = ?", req.ID).UpdateColumn("is_read", req.IsRead).Error
}

func (p NotifyStore) GetNotifyPaging(ctx context.Context, param request.ListNotifyReq) (pagination api_response.Pagination, err error) {
	pagination.Limit = param.Limit
	pagination.Page = param.Page
	var place []*entity.Notify

	exec := p.DbExecutor.Scopes()

	exec.
		Scopes(api_response.Paginate(place, &pagination, exec)).
		Preload("FromUser").
		Preload("ToUser").
		Where("to_user_id = ?", param.UserID).
		Order("created_at desc").
		Find(&place)

	pagination.Rows = place
	return
}

func (p NotifyStore) CountNotifyRead(ctx context.Context, userID string) (int64, error) {
	var count int64
	err := p.DbExecutor.Model(&entity.Notify{}).
		Where("to_user_id = ? AND is_read = false", userID).Count(&count)
	if err.Error != nil {
		return count, err.Error
	}
	return count, nil
}
