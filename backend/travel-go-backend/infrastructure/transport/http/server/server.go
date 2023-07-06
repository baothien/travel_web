package server

import (
	"context"
	"flag"
	"fmt"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger/level"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/route"
	"gitlab.com/virtual-travel/travel-go-backend/utils/i18nutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

const (
	MB = 1024 * 1024
)

var (
	logLevel = flag.String("log-level", "debug", "HTTP Server log level")
	//configFile = flag.String("config-file", "", "HTTP Server config file") /
)

func initFlag() {
	flag.Parse()
}

func initLogger() {
	logger.NewLogger(level.Level(*logLevel))
}

func initI18NBundle() {
	i18nutil.NewBundle()
}

func init() {
	initFlag()
	initLogger()
	initI18NBundle()
}

type ServerOption func(*HTTPServer)

type HTTPServer struct {
	Name                    string
	Port                    int
	StrictSlash             bool
	Routes                  []route.Route
	GroupRoutes             []route.GroupRoute
	Middlewares             []func(c *gin.Context)
	GinOptions              []route.GinOption
	GracefulShutdownTimeout time.Duration
	OnCloseFunc             func()
}

func NewHTTPServer(options ...ServerOption) *HTTPServer {
	s := &HTTPServer{}
	for _, option := range options {
		option(s)
	}
	return s
}

func (s *HTTPServer) Run(r *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	//r := route.NewGin(
	//	route.AddMiddlewares(s.Middlewares...),
	//	route.AddHealthCheckRoute(),
	//	route.StrictSlash(s.StrictSlash),
	//	route.AddGroupRoutes(s.GroupRoutes),
	//	route.AddRoutes(s.Routes),
	//	route.AddGinOptions(s.GinOptions...))
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	hs := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: r,
	}

	// Graceful shutdown
	idleConnectionClosed := make(chan struct{})
	go func() {
		cs := make(chan os.Signal, 1)
		signal.Notify(cs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		<-cs

		ctx, cancel := context.WithTimeout(context.Background(), s.GracefulShutdownTimeout)
		defer cancel()

		logger.Info("Server is shuting down")

		if err := hs.Shutdown(ctx); err != nil {
			logger.Error("Server has failed graceful shutdown", zap.String("error", err.Error()))
		}

		s.OnCloseFunc()

		<-ctx.Done()
		close(idleConnectionClosed)
	}()

	go func() {
		logger.Infof("Starting server: %s with Insecure, and listen at port: %d", s.Name, s.Port)
		if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server has failed run", zap.String("error", err.Error()))
			os.Exit(1)
		} else {
			logger.Info("Server has graceful shutdown completely")
		}
	}()

	<-idleConnectionClosed
}

func AddName(n string) ServerOption {
	return func(s *HTTPServer) {
		s.Name = n
	}
}

func AddPort(p int) ServerOption {
	return func(s *HTTPServer) {
		s.Port = p
	}
}

func AddMiddlewares(m []func(c *gin.Context)) ServerOption {
	return func(s *HTTPServer) {
		s.Middlewares = append(s.Middlewares, m...)
	}
}

func AddOnCloseFunc(f func()) ServerOption {
	return func(s *HTTPServer) {
		s.OnCloseFunc = f
	}
}

func AddGinOptions(o []route.GinOption) ServerOption {
	return func(s *HTTPServer) {
		s.GinOptions = o
	}
}

func (s *HTTPServer) AddRoutes(r []route.Route) {
	s.Routes = r
}

func (s *HTTPServer) AddGroupRoutes(gr []route.GroupRoute) {
	s.GroupRoutes = gr
}

func StrictSlash() ServerOption {
	return func(s *HTTPServer) {
		s.StrictSlash = true
	}
}

func SetGracefulShutdownTimeout(d time.Duration) ServerOption {
	return func(s *HTTPServer) {
		s.GracefulShutdownTimeout = d
	}
}

func MustRun(s *HTTPServer, err error, r *gin.Engine) {
	fmt.Println("ELASTIC_APM_SERVICE_NAME = ", os.Getenv("ELASTIC_APM_SERVICE_NAME"))
	if err != nil {
		panic(err)
	}
	s.Run(r)
}
