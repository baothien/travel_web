package http

import (
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/server"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/config"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/delivery/http/handler"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/repository"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/usecase"
)

type Serve struct {
	AccountHandler handler.AccountHandler
}

func NewServer(config config.Config, db pg.Database) (*server.HTTPServer, error) {
	s := server.NewHTTPServer(
		server.AddName(config.ServiceName),
		server.AddPort(config.ServerPort),
		server.StrictSlash(),
	)

	srv := Serve{
		AccountHandler: handler.NewAccountHandler(
			usecase.NewAccountUseCase(repository.NewAccountStoreImpl(db)),
		),
	}

	// init router for account handler
	s.AddGroupRoutes(AccountRoutes(srv))
	return s, nil
}
