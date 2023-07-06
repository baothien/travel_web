package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/method"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/route"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/middlewares"
)

func NotifyRoutes(serve Serve) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/ws",
			Routes: []route.Route{
				{
					Path:    "/:id",
					Method:  method.GET,
					Handler: serve.NotifyHandler.SocketNotify,
				},
			},
		},
		{
			Prefix: "/notify",
			Routes: []route.Route{
				{
					Path:   "/create",
					Method: method.POST,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.NotifyHandler.CreateNotify,
				},
				{
					Path:   "/list",
					Method: method.GET,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.NotifyHandler.GetNotifyPaging,
				},
				{
					Path:   "/count",
					Method: method.GET,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.NotifyHandler.GetCountNotify,
				},
				{
					Path:    "/push",
					Method:  method.POST,
					Handler: serve.NotifyHandler.PushNotifySocket,
				},
				{
					Path:   "/is-read/:id",
					Method: method.PATCH,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.NotifyHandler.UpdateIsReadNotify,
				},
			},
		},
	}
}
