package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	pg "gitlab.com/virtual-travel/travel-go-backend/infrastructure/database"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/route"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/transport/http/server"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/config"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/delivery/http"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/docs"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/migrations"
	"gitlab.com/virtual-travel/travel-go-backend/utils/validatorutil"

	"go.uber.org/zap"
	"time"
)

func main() {
	appConfig, err := loadAppConf()
	if err != nil {
		logger.Error("can not load app config")
		panic(err)
	}

	db, err := connectDatabase(appConfig)
	if err != nil {
		logger.Error("can not connect database")
		panic(err)
	}

	// migrate db..
	errMigrate := migrations.AutoMigrateDB(db)
	if errMigrate != nil {
		logger.Error("Error auto migrate db")
		panic(errMigrate)
	}

	httpServer, err := http.NewServer(appConfig, *db)

	//add gin.
	r := route.NewGin(
		route.AddMiddlewares(httpServer.Middlewares...),
		route.AddHealthCheckRoute(),
		route.StrictSlash(httpServer.StrictSlash),
		route.AddGroupRoutes(httpServer.GroupRoutes),
		route.AddRoutes(httpServer.Routes),
		route.AddGinOptions(httpServer.GinOptions...))

	binding.Validator = new(validatorutil.DefaultValidator)

	// add docs
	docs.SwaggerInfo.BasePath = fmt.Sprintf(
		"/%v/%v",
		viper.Get("BUILD_ENV").(string),
		viper.Get("SERVICE_NAME").(string))
	docs.SwaggerInfo.Title = fmt.Sprintf("This is: %v service", viper.Get("SERVICE_NAME").(string))
	docs.SwaggerInfo.Description = fmt.Sprintf("Travel")
	docs.SwaggerInfo.Version = fmt.Sprintf("v1.0.0")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Static("/files", "../../public/files")

	server.MustRun(httpServer, err, r)
}

func loadAppConf() (config.Config, error) {
	appConfig, err := config.LoadConfig("../../config/", "app")
	if err != nil {
		logger.Error("cannot load config", zap.String("errors", err.Error()))
	}
	return appConfig, err
}

func connectDatabase(config config.Config) (*pg.Database, error) {
	chainConnectionString := []pg.ConnectionOption{
		pg.SetConnection(config.DBHost, config.DBPort),
		pg.SetLoginCredentials(config.DBUsername, config.DBPassword),
		pg.SetDatabase(config.DBName),
		pg.SetConnectionTimeout(time.Second * 5),
		pg.SetMaxOpenConnections(1000),
		pg.SetMaxIdleConnections(100),
		pg.SetConnectionMaxIdleTime(15 * time.Second),
		pg.SetConnectionMaxLifeTime(1 * time.Hour),
		pg.SetFallbackConnection("", ""),
		pg.SetSSL(pg.Disable, "", "", ""),
	}
	db, err := pg.NewDB(
		"gorm",
		pg.AddChainConnectionOptions(chainConnectionString...),
		pg.SetLogLevel(pg.Debug),
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
