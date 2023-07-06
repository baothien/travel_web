package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/method"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/route"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/middlewares"
)

func PlaceTypeRoutes(serve Serve) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/place-type",
			Routes: []route.Route{
				{
					Path:    "/create",
					Method:  method.POST,
					Handler: serve.PlaceTypeHandler.CreatePlaceType,
				},
				{
					Path:    "/list",
					Method:  method.GET,
					Handler: serve.PlaceTypeHandler.GetPlaceType,
				},
			},
		},
		{
			Prefix: "/place",
			Routes: []route.Route{
				{
					Path:    "/create",
					Method:  method.POST,
					Handler: serve.PlaceHandler.CreatePlace,
				},
				{
					Path:    "/update/:id",
					Method:  method.PATCH,
					Handler: serve.PlaceHandler.UpdatePlace,
				},
				{
					Path:    "/list",
					Method:  method.GET,
					Handler: serve.PlaceHandler.GetPlace,
				},
				{
					Path:    "/detail/:id",
					Method:  method.GET,
					Handler: serve.PlaceHandler.GetDetailPlace,
				},
				{
					Path:   "/favorite",
					Method: method.POST,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.PlaceHandler.AddFavoritePlace,
				},
				{
					Path:   "/favorite/list",
					Method: method.GET,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.PlaceHandler.GetFavoritePlace,
				},
			},
		},
		{
			Prefix: "/review",
			Routes: []route.Route{
				{
					Path:   "/create",
					Method: method.POST,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.PlaceHandler.CreateReviewPlace,
				},
				{
					Path:    "/list/:place_id",
					Method:  method.GET,
					Handler: serve.PlaceHandler.GetReviewPlace,
				},
				{
					Path:   "/child/create/:parent_id",
					Method: method.POST,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.PlaceHandler.CreateChildReviewPlace,
				},
			},
		},
		{
			Prefix: "/child-review",
			Routes: []route.Route{
				{
					Path:    "/list/:review_id",
					Method:  method.GET,
					Handler: serve.PlaceHandler.GetChildReview,
				},
			},
		},
		{
			Prefix: "/cms",
			Routes: []route.Route{
				{
					Path:    "/place/list",
					Method:  method.GET,
					Handler: serve.PlaceHandler.GetPlacePaging,
				},
			},
		},
		{
			Prefix: "/checkin",
			Routes: []route.Route{
				{
					Path:   "/create",
					Method: method.POST,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.CheckinHandler.CreateCheckin,
				},
				{
					Path:   "/list",
					Method: method.GET,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.CheckinHandler.GetCheckin,
				},
			},
		},
	}
}
