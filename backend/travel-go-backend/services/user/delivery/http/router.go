package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/method"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/route"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/middlewares"
)

func AccountRoutes(serve Serve) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/user",
			Routes: []route.Route{
				{
					Path:    "/register",
					Method:  method.POST,
					Handler: serve.AccountHandler.RegisterUser,
				},
				{
					Path:    "/login",
					Method:  method.POST,
					Handler: serve.AccountHandler.Login,
				},
				{
					Path:   "/profile",
					Method: method.GET,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.AccountHandler.Profile,
				},
				{
					Path:   "/profile",
					Method: method.PATCH,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.AccountHandler.UpdateProfile,
				},
				{
					Path:   "/list",
					Method: method.GET,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.AccountHandler.ListAccount,
				},
				{
					Path:   "/detail/:id",
					Method: method.GET,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.AccountHandler.DetailUser,
				},
				{
					Path:   "/refresh-token",
					Method: method.GET,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTRefreshMiddleware().MiddlewareFunc(),
					},
					Handler: serve.AccountHandler.RefreshToken,
				},
				{
					Path:   "/change-pass",
					Method: method.PATCH,
					Middlewares: []func(c *gin.Context){
						middlewares.TrackRequest,
						middlewares.JWTMiddleware().MiddlewareFunc(),
					},
					Handler: serve.AccountHandler.ChangePassword,
				},
			},
		},
	}
}
