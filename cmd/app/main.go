package main

import (
	"context"
	"flag"
	"fmt"
	lg "log"
	"os"
	"os/signal"
	"syscall"
	"time"

	cfg "gitlab.geogracom.com/skdf/skdf-abac-go/configs/app"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/handlers"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/servers/http"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/services"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/auth"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/logger"
)

const (
	casbinPath = "configs/app/casbin/"
)

func main() {
	modePtr := flag.String("mode", "debug", "mode defines which env file to use")
	flag.Parse()

	// Configurations
	var config cfg.Config

	switch *modePtr {
	case "debug":
		config = cfg.Load("configs/app/.env.debug")
	case "release":
		config = cfg.Load("configs/app/.env")
	default:
		lg.Fatalf("invalid mode %s, please check mode is valid", *modePtr)
	}

	// Logger
	log := logger.New("debug", "skdf-abac-go-logger")
	defer func() {
		err := logger.CleanUp(log)
		log.Fatal("failed to cleanup logs", logger.Error(err))
	}()

	enforcer, err := auth.NewEnforcer(casbinPath+"model.conf", casbinPath+"policy.csv")
	if err != nil {
		log.Fatal("failed to initialize auth enforcer", logger.Error(err))
	}

	start := time.Now()
	ctx := context.Background()
	_db, err := db.ConnectContext(ctx, config.DB)
	if err != nil {
		log.Fatal("failed to initialize db", logger.Error(err))
	}
	fmt.Println("connection to database took: ", time.Since(start))

	repo := repositories.NewRepository(_db, enforcer)

	service := services.NewService(repo)

	router := handlers.NewRouter(handlers.Option{
		Conf:    config,
		Logger:  log,
		Service: service,
	})

	server := http.NewServer(config.HTTP.Port, router)

	go func() {
		if err = server.Run(); err != nil {
			log.Fatal("failed to run http server", logger.Error(err))
			panic(err)
		}
	}()

	log.Info("Server started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("failed to stop server", logger.Error(err))
	}

}
