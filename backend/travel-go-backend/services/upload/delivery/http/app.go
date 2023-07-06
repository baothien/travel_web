package http

import (
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/server"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/config"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/delivery/http/handler"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/domain/places/repository"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/domain/places/usecase"
)

type Serve struct {
	FileHandler handler.FileHandler
}

func NewServer(config config.Config, db pg.Database) (*server.HTTPServer, error) {
	s := server.NewHTTPServer(
		server.AddName(config.ServiceName),
		server.AddPort(config.ServerPort),
		server.StrictSlash(),
	)

	srv := Serve{
		FileHandler: handler.NewFileHandler(
			usecase.NewFileUseCase(repository.NewFileStoreImpl(db)),
		),
	}

	// init router for account handler
	s.AddGroupRoutes(FileRoutes(srv))
	return s, nil
}
