// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/David-Kalashir/crs-server/config"
	v1 "github.com/David-Kalashir/crs-server/internal/controller/http/v1"
	"github.com/David-Kalashir/crs-server/internal/usecase"
	"github.com/David-Kalashir/crs-server/internal/usecase/repo"
	"github.com/David-Kalashir/crs-server/pkg/httpserver"
	"github.com/David-Kalashir/crs-server/pkg/logger"
	"github.com/David-Kalashir/crs-server/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	loginUseCase := usecase.New(
		repo.New(pg),
	)
	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, loginUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
