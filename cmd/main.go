package main

import (
	"fmt"

	"github.com/mado/config"
	_ "github.com/mado/docs"
	"github.com/mado/service"
	"github.com/mado/storage"

	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mado/storage/postgre"
	cache "github.com/mado/storage/redis"
	server "github.com/mado/transport/http"
	"github.com/mado/transport/http/handler"
	"go.uber.org/zap"
)

// @title 	ЦАРКА REST API
// @version	1.0
// @description ЦАРКА Тествое задание.
// @termsOfService	http://swagger.io/terms/
// @host localhost:8000
// @BasePath	/rest
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// load env
	var once sync.Once
	once.Do(config.PrepareENV)

	// get config
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	// init logger

	l, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	defer func(*zap.Logger) {
		err := l.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(l)

	// init repo layer
	db, err := postgre.Dial(constructDSN(cfg))
	if err != nil {
		return err
	}
	defer db.Close()

	redisClient, err := cache.NewRedisClient(cfg)
	if err != nil {
		return err
	}

	repo := storage.NewStorage(db, redisClient)

	// init service layer
	usecases := service.NewService(repo, l)

	// init handler layer
	controller := handler.NewHandler(usecases, l)

	// init http server instance
	httpServer := server.NewServer(cfg, controller)

	l.Info("Start app", zap.String("port", cfg.AppPort))
	httpServer.StartServer()

	// grace full shutdown
	osSignCh := make(chan os.Signal, 1)
	signal.Notify(osSignCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-osSignCh:
		l.Info("signal accepted: ", zap.String("signal", s.String()))
	case err = <-httpServer.Notify:
		l.Info("server closing", zap.Error(err))
	}

	if err = httpServer.Shutdown(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("error while shutting down server: %s", err)
	}

	return nil
}

func constructDSN(cfg *config.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.TZ)
}
