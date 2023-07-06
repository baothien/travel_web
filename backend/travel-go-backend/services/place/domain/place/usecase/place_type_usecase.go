package usecase

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/repository"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/constant"
	"gitlab.com/virtual-travel/travel-go-backend/utils/randutil"
	"go.uber.org/zap"
)

type PlaceTypeUseCase struct {
	PlaceTypeRepository repository.PlaceTypeRepository
}

func NewPlaceTypeUseCase(accountRepository repository.PlaceTypeRepository) PlaceTypeUseCase {
	return PlaceTypeUseCase{
		PlaceTypeRepository: accountRepository,
	}
}

func (p PlaceTypeUseCase) CreatePlaceType(ctx context.Context, req request.CreatePlaceTypeReq) error {
	placeType := entity.PlaceType{
		ID:   randutil.UUIDRand(),
		Code: req.Code,
		Name: req.Name,
	}

	err := p.PlaceTypeRepository.CreatePlaceType(ctx, placeType)
	if err != nil {
		logger.Error("Error create place type", zap.Error(err))
		return errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	return nil
}

func (p PlaceTypeUseCase) GetPlaceType(ctx context.Context) (places []entity.PlaceType, err error) {
	places, err = p.PlaceTypeRepository.GetPlaceType(ctx)
	if err != nil {
		return places, errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	return
}
