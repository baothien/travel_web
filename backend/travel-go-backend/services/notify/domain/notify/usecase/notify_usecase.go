package usecase

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/domain/notify/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/domain/notify/repository"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/constant"
	"go.uber.org/zap"
)

type NotifyUseCase struct {
	NotifyRepository repository.NotifyRepository
}

func NewNotifyUseCase(accountRepository repository.NotifyRepository) NotifyUseCase {
	return NotifyUseCase{
		NotifyRepository: accountRepository,
	}
}

func (n NotifyUseCase) CreateNotify(ctx context.Context, place entity.Notify) error {
	err := n.NotifyRepository.CreateNotify(ctx, place)
	if err != nil {
		logger.Error("Error create notify", zap.Error(err))
		return err
	}
	return nil
}

func (n NotifyUseCase) GetNotifyPaging(ctx context.Context, param request.ListNotifyReq) (api_response.Pagination, error) {
	places, err := n.NotifyRepository.
		GetNotifyPaging(ctx, param)
	if err != nil {
		return places, errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	return places, nil
}

func (n NotifyUseCase) UpdateIsRead(ctx context.Context, req request.UpdateIsRead) error {
	err := n.NotifyRepository.UpdateIsRead(ctx, req)
	if err != nil {
		logger.Error("Error update is read", zap.Error(err))
		return errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	return nil
}

func (n NotifyUseCase) CountNotifyRead(ctx context.Context, userID string) (int64, error) {
	count, err := n.NotifyRepository.
		CountNotifyRead(ctx, userID)
	if err != nil {
		return count, errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	return count, nil
}
