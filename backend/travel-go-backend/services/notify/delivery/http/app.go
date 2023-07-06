package http

import (
	"flag"
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/server"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/config"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/delivery/http/handler"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/domain/notify/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/domain/notify/repository"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/domain/notify/usecase"
)

type Serve struct {
	NotifyHandler handler.NotifyHandler
}

func NewServer(config config.Config, db pg.Database) (*server.HTTPServer, error) {
	s := server.NewHTTPServer(
		server.AddName(config.ServiceName),
		server.AddPort(config.ServerPort),
		server.StrictSlash(),
	)

	flag.Parse()
	hub := entity.NewHub()

	go hub.Run()

	srv := Serve{
		NotifyHandler: handler.NewNotifyHandler(
			usecase.NewNotifyUseCase(repository.NewNotifyStoreImpl(db)),
			*hub,
		),
	}

	// init router for account handler
	s.AddGroupRoutes(NotifyRoutes(srv))
	return s, nil
}
