package repository

import (
	"context"
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/entity"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, user entity.User) error
	UpdateAccount(ctx context.Context, user entity.User) error
	GetUserByUseName(ctx context.Context, userName string) (entity.User, error)
	GetUserByUserId(ctx context.Context, id string) (entity.User, error)
	ListUser(ctx context.Context, req request.ListUser) (api_response.Pagination, error)
}

type AccountStore struct {
	DbExecutor *gorm.DB
}

func NewAccountStoreImpl(db pg.Database) AccountRepository {
	return AccountStore{
		DbExecutor: db.Gorm,
	}
}

func (a AccountStore) CreateAccount(ctx context.Context, user entity.User) error {
	return a.DbExecutor.Create(&user).Error
}

func (a AccountStore) UpdateAccount(ctx context.Context, user entity.User) error {
	return a.DbExecutor.Updates(&user).Error
}

func (a AccountStore) GetUserByUseName(ctx context.Context, userName string) (user entity.User, err error) {
	err = a.DbExecutor.Where("user_name = ?", userName).First(&user).Error
	return
}

func (a AccountStore) GetUserByUserId(ctx context.Context, id string) (user entity.User, err error) {
	err = a.DbExecutor.Where("id = ?", id).First(&user).Error
	return
}

func (a AccountStore) ListUser(ctx context.Context, param request.ListUser) (pagination api_response.Pagination, err error) {
	pagination.Limit = param.Limit
	pagination.Page = param.Page
	var user []*entity.User

	exec := a.DbExecutor.Scopes()
	//
	exec.
		Scopes(api_response.Paginate(user, &pagination, exec)).
		Omit("password").
		Order("created_at desc").
		Find(&user)

	pagination.Rows = user

	return pagination, nil

}
