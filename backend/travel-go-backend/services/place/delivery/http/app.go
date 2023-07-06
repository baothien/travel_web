package http

import (
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/server"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/config"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/delivery/http/handler"
	checkinrp "gitlab.com/virtual-travel/travel-go-backend/services/place/domain/checkin/repository"
	checkinus "gitlab.com/virtual-travel/travel-go-backend/services/place/domain/checkin/usecase"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/repository"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/usecase"
	reviewrp "gitlab.com/virtual-travel/travel-go-backend/services/place/domain/review/repository"
	reviewus "gitlab.com/virtual-travel/travel-go-backend/services/place/domain/review/usecase"
)

type Serve struct {
	PlaceTypeHandler handler.PlaceTypeHandler
	PlaceHandler     handler.PlaceHandler
	CheckinHandler   handler.CheckinHandler
}

func NewServer(config config.Config, db pg.Database) (*server.HTTPServer, error) {
	s := server.NewHTTPServer(
		server.AddName(config.ServiceName),
		server.AddPort(config.ServerPort),
		server.StrictSlash(),
	)

	srv := Serve{
		PlaceTypeHandler: handler.NewPlaceTypeHandler(
			usecase.NewPlaceTypeUseCase(repository.NewPlaceTypeStoreImpl(db)),
		),
		PlaceHandler: handler.NewPlaceHandler(
			usecase.NewPlaceUseCase(repository.NewPlaceStoreImpl(db)),
			reviewus.NewReviewUseCase(reviewrp.NewReviewStoreImpl(db), repository.NewPlaceStoreImpl(db)),
		),
		CheckinHandler: handler.NewCheckinHandler(
			checkinus.NewCheckinUseCase(checkinrp.NewCheckinStoreImpl(db)),
		),
	}

	// init router for account handler
	s.AddGroupRoutes(PlaceTypeRoutes(srv))
	return s, nil
}
