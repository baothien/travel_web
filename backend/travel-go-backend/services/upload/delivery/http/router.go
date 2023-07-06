package http

import (
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/method"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/route"
)

func FileRoutes(serve Serve) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/file",
			Routes: []route.Route{
				{
					Path:    "/upload",
					Method:  method.POST,
					Handler: serve.FileHandler.UploadFile,
				},
			},
		},
	}
}
