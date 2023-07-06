package repository

import (
	"context"
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/domain/places/entity"
	"gorm.io/gorm"
)

type FileRepository interface {
	CreateFile(ctx context.Context, file entity.File) error
}

type FileStore struct {
	DbExecutor *gorm.DB
}

func NewFileStoreImpl(db pg.Database) FileRepository {
	return FileStore{
		DbExecutor: db.Gorm,
	}
}

func (f FileStore) CreateFile(ctx context.Context, file entity.File) error {
	err := f.DbExecutor.Create(&file)
	return err.Error
}
