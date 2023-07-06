package pg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/pkg/errors"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pgxv4"
)

type Database struct {
	Sqlx *sqlx.DB
	Gorm *gorm.DB
}

func NewDB(lib string, options ...ConnectionOption) (*Database, error) {
	cfg := Connection{
		SSLMode:               Disable,
		LogLevel:              Debug,
		MaxOpenConnections:    100,
		MaxIdleConnections:    25,
		ConnectionMaxIdleTime: 5 * time.Minute,
		ConnectionMaxLifeTime: 10 * time.Minute,
		ConnectionTimeout:     30 * time.Second,
		ConnectionOptions:     []ConnectionOption{},
	}

	for _, option := range options {
		option(&cfg)
	}

	// Chain Connection Options
	if len(cfg.ConnectionOptions) > 0 {
		for _, option := range cfg.ConnectionOptions {
			option(&cfg)
		}
	}

	pgConfig, err := pgx.ParseConfig(cfg.ToConnectionString())
	if err != nil {
		return nil, errors.Wrap(err, "call pgx.ParseConfig() has failed")
	}

	db, err := apmsql.Open("pgx", pgConfig.ConnString())
	if err != nil {
		return nil, errors.Wrap(err, "call apmsql.Open() has failed")
	}

	dbx := sqlx.NewDb(db, "pgx")
	if err := dbx.Ping(); err != nil {
		return nil, errors.Wrap(err, "call db.Ping has failed")
	}
	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetConnMaxIdleTime(cfg.ConnectionMaxIdleTime)
	db.SetConnMaxLifetime(cfg.ConnectionMaxLifeTime)

	dbx.Mapper = reflectx.NewMapperTagFunc("json", nil, strings.ToLower)
	if lib == "gorm" {
		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			Conn: dbx,
		}), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             time.Millisecond * 200,
					LogLevel:                  logger.Warn,
					IgnoreRecordNotFoundError: false,
					Colorful:                  true,
				},
			),
			SkipDefaultTransaction: true,
			AllowGlobalUpdate:      false,
		})

		if err != nil {
			return nil, errors.Wrap(err, "call apmsql.Open() has failed")
		}

		return &Database{
			Gorm: gormDB,
			Sqlx: dbx,
		}, nil
	}

	return &Database{
		Sqlx: dbx,
	}, nil
}
