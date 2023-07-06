package usecase

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/domain/places/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/domain/places/repository"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/constant"
	"gitlab.com/virtual-travel/travel-go-backend/utils/fileutil"
	"gitlab.com/virtual-travel/travel-go-backend/utils/randutil"

	"go.uber.org/zap"
	"io"
	"os"
)

type FileUseCase struct {
	FileRepository repository.FileRepository
}

func NewFileUseCase(accountRepository repository.FileRepository) FileUseCase {
	return FileUseCase{
		FileRepository: accountRepository,
	}
}

func (u FileUseCase) UploadFile(ctx context.Context, body request.FileReq) (entity.File, error) {
	fmt.Println(1)
	// Source
	src, err := body.File.Open()
	if err != nil {
		logger.Error("Error open file", zap.Error(err))
		return entity.File{}, errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	defer src.Close()
	fmt.Println(2)

	_, savePath, filePath := fileutil.MoveFileUploads(body.Type, body.File.Header.Get("Content-Type"))
	// Destination
	dst, err := os.Create(savePath)
	if err != nil {
		return entity.File{}, errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	defer dst.Close()
	fmt.Println(3)

	if _, err = io.Copy(dst, src); err != nil {
		logger.Error("Error copy file", zap.Error(err))
		return entity.File{}, errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	fmt.Println(4)

	// init file struct
	file := entity.File{
		ID:          randutil.UUIDRand(),
		Name:        body.File.Filename,
		Size:        uint64(body.File.Size),
		Type:        body.Type,
		ContentType: body.File.Header.Get("Content-Type"),
		Path:        filePath,
		FullPath: fmt.Sprintf(
			"%v/%v/%v/%v",
			viper.Get("BASE_URL_SERVICE").(string),
			viper.Get("BUILD_ENV").(string),
			viper.Get("SERVICE_NAME").(string),
			filePath),
	}

	fmt.Println(5)
	err = u.FileRepository.CreateFile(ctx, file)
	if err != nil {
		logger.Error("Error save file to db", zap.Error(err))
		return file, errors.New(constant.ERROR_INTERNAL_SERVER)
	}
	fmt.Println(6)

	return file, nil
}
