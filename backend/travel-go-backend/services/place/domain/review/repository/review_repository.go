package repository

import (
	"context"
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/review/entity"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	CreateReview(ctx context.Context, review entity.Review) error
	CreateChildReview(ctx context.Context, childReview entity.ChildReview) error
	GetReviewByID(ctx context.Context, id string) (entity.Review, error)
	GetReviewByPlaceID(ctx context.Context, param request.ListReviewPlaceReq) (pagination api_response.Pagination, err error)
	GetChildReviewByReviewID(ctx context.Context, param request.ListChildReviewReq) (pagination api_response.Pagination, err error)
}

type ReviewStore struct {
	DbExecutor *gorm.DB
}

func NewReviewStoreImpl(db pg.Database) ReviewRepository {
	return ReviewStore{
		DbExecutor: db.Gorm,
	}
}

func (r ReviewStore) CreateReview(ctx context.Context, review entity.Review) error {
	return r.DbExecutor.Create(&review).Error
}

func (r ReviewStore) CreateChildReview(ctx context.Context, childReview entity.ChildReview) error {
	return r.DbExecutor.Create(&childReview).Error
}

func (r ReviewStore) GetReviewByID(ctx context.Context, id string) (entity.Review, error) {
	review := entity.Review{}
	err := r.DbExecutor.
		Where("id = ?", id).
		Preload("ChildReview").
		First(&review).Error
	if err != nil {
		return review, err
	}
	return review, nil
}

func (r ReviewStore) GetReviewByPlaceID(ctx context.Context, param request.ListReviewPlaceReq) (pagination api_response.Pagination, err error) {
	var review []entity.Review

	pagination.Limit = param.Limit
	pagination.Page = param.Page

	exec := r.DbExecutor.Scopes()

	exec.
		Scopes(api_response.Paginate(review, &pagination, exec)).
		Preload("ReviewImg").
		Preload("ChildReview").
		Preload("ChildReview.User").
		Preload("ChildReview.ReviewImg").
		Preload("User").
		Order("created_at desc").
		Where("place_id = ?", param.PlaceID).
		Find(&review)

	pagination.Rows = review

	return pagination, nil
}

func (r ReviewStore) GetChildReviewByReviewID(ctx context.Context, param request.ListChildReviewReq) (pagination api_response.Pagination, err error) {
	var review []entity.ChildReview

	pagination.Limit = param.Limit
	pagination.Page = param.Page

	exec := r.DbExecutor.Scopes()

	exec.
		Scopes(api_response.Paginate(review, &pagination, exec)).
		Preload("ReviewImg").
		Preload("User").
		Order("created_at desc").
		Where("parent_id = ?", param.ReviewID).
		Find(&review)

	pagination.Rows = review

	return pagination, nil
}
