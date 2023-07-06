package usecase

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/repository"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/constant"
	"gitlab.com/virtual-travel/travel-go-backend/utils/randutil"
	"go.uber.org/zap"
	"time"
)

type PlaceUseCase struct {
	PlaceRepository repository.PlaceRepository
}

func NewPlaceUseCase(accountRepository repository.PlaceRepository) PlaceUseCase {
	return PlaceUseCase{
		PlaceRepository: accountRepository,
	}
}

func (p PlaceUseCase) CreatePlace(ctx context.Context, req request.CreatePlaceReq) error {
	placeId := randutil.UUIDRand()
	placeImg := []entity.PlaceImg{}

	for i := range req.Img {
		placeImgItem := entity.PlaceImg{
			ID:      randutil.UUIDRand(),
			PlaceID: placeId,
			Name:    req.Img[i].Name,
			Url:     req.Img[i].Url,
		}

		placeImg = append(placeImg, placeImgItem)
	}

	placeImgCheckin := []entity.PlaceImgCheckin{}

	for i := range req.Img {
		placeImgCheckinItem := entity.PlaceImgCheckin{
			ID:      randutil.UUIDRand(),
			PlaceID: placeId,
			Name:    req.ImgCheckin[i].Name,
			Url:     req.ImgCheckin[i].Url,
		}

		placeImgCheckin = append(placeImgCheckin, placeImgCheckinItem)
	}

	place := entity.Place{
		ID:              placeId,
		Thumbnail:       req.Thumbnail,
		Name:            req.Name,
		PlaceTypeID:     &req.PlaceTypeID,
		Lat:             req.Lat,
		Lng:             req.Lng,
		Address:         req.Address,
		PlaceImg:        placeImg,
		PlaceImgCheckin: placeImgCheckin,
		Review:          nil,
		UserFavorite:    nil,
		PlaceType:       nil,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
	}

	err := p.PlaceRepository.CreatePlace(ctx, place)
	if err != nil {
		logger.Error("Error create place", zap.Error(err))
		return errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	return nil
}

func (p PlaceUseCase) UpdatePlace(ctx context.Context, place entity.Place) error {
	err := p.PlaceRepository.UpdatePlace(ctx, place)
	if err != nil {
		logger.Error("Error update place", zap.Error(err))
		return errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	return nil
}

func (p PlaceUseCase) GetPlace(ctx context.Context) (places []entity.Place, err error) {
	places, err = p.PlaceRepository.GetPlace(ctx)
	if err != nil {
		return places, errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	return
}

func (p PlaceUseCase) GetDetailPlace(ctx context.Context, id string) (entity.Place, error) {
	place, err := p.PlaceRepository.GetPlaceById(ctx, id)
	if err != nil {
		return place, errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	return place, nil
}

func (p PlaceUseCase) GetPlacePaging(ctx context.Context, param request.ListPlaceReq) (api_response.Pagination, error) {
	places, err := p.PlaceRepository.GetPlacePaging(ctx, param)
	if err != nil {
		return places, errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	return places, nil
}

func (p PlaceUseCase) FavoritePlace(ctx context.Context, param request.FavoritePlaceAddReq) error {
	err := p.PlaceRepository.FavoritePlace(ctx, param)
	if err != nil {
		logger.Error("error add favorite place", zap.Error(err))
		return errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	return nil
}

func (p PlaceUseCase) GetFavoritePlacePaging(ctx context.Context, param request.ListFavoritePlaceReq) (api_response.Pagination, error) {
	places, err := p.PlaceRepository.GetFavoritePlacePaging(ctx, param)
	if err != nil {
		return places, errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	return places, nil
}
