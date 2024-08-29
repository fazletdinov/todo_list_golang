package app

import (
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"todo-list/config"
	"todo-list/internal/database"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Application struct {
	DB    *gorm.DB
	Env   *config.Config
	Redis *redis.Client
	Log   *zerolog.Logger
}

func App() Application {
	app := &Application{}
	Env, errEnv := config.InitConfig()
	if errEnv != nil {
		log.Fatalf("ошибка загрузки ENV - %v", errEnv)
	}
	PostgresClient, errPostgres := database.InitDatabse(Env)
	if errPostgres != nil {
		log.Fatalf("ошибка подключения к Postgres - %v", errPostgres)
	}
	RedisClient, errRedis := database.InitRedisDB(Env)
	if errRedis != nil {
		log.Fatalf("ошибка подключения к Redis - %v", errPostgres)
	}

	log := setupLogger(Env.Env)
	app.Env = Env
	app.DB = PostgresClient
	app.Redis = RedisClient
	app.Log = log
	return *app
}

func setupLogger(env string) *zerolog.Logger {
	zerolog.TimeFieldFormat = "02/Jan/2006 - 15:04:05 -0700"
	switch env {
	case envLocal:
		logger := zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "02/Jan/2006 - 15:04:05 -0700",
		}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()
		return &logger
	case envDev:
		logger := zerolog.New(os.Stdout).
			Level(zerolog.DebugLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()
		return &logger
	case envProd:
		logger := zerolog.New(os.Stdout).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()
		return &logger
	}
	return nil
}
