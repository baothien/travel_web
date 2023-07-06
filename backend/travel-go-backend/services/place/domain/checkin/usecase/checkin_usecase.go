package usecase

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/checkin/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/checkin/repository"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/constant"
	"gitlab.com/virtual-travel/travel-go-backend/utils/randutil"
	"go.uber.org/zap"
)

type CheckinUseCase struct {
	CheckinRepository repository.CheckinRepository
}

func NewCheckinUseCase(accountRepository repository.CheckinRepository) CheckinUseCase {
	return CheckinUseCase{
		CheckinRepository: accountRepository,
	}
}

func (c CheckinUseCase) CreateCheckin(ctx context.Context, checkin entity.Checkin) error {
	checkinId := randutil.UUIDRand()
	checkin.ID = checkinId

	err := c.CheckinRepository.CreateCheckin(ctx, checkin)
	if err != nil {
		logger.Error("Error create checkin", zap.Error(err))
		return errors.New(constant.ERROR_INTERNAL_SERVER)
	}

	return nil
}

func (c CheckinUseCase) GetCheckinPaging(ctx context.Context, param request.ListCheckinReq) (api_response.Pagination, error) {
	checkins, err := c.CheckinRepository.ListCheckin(ctx, param)
	if err != nil {
		return checkins, errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	return checkins, nil
}
