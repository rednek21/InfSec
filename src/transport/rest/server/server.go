package server

import (
	"InfSec/pkg/logger"
	"InfSec/src/config"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Server struct {
	Host string
	Port int
	log  logger.ILogger
}

func New(cfg *config.Config, logger logger.ILogger) *Server {
	return &Server{
		Host: cfg.Server.Host,
		Port: cfg.Server.Port,
		log:  logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	GSCtx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	router.Use(gin.Recovery())

	setupRoutes(router)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(s.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Fatal("Failed to listen and serve: %v\n", err)
		}
	}()
	s.log.Info("Server started on :%d", s.Port)

	<-GSCtx.Done()

	stop()
	s.log.Info("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		s.log.Fatal("Server forced to shutdown: ", err)
	}
	s.log.Info("Server exiting")

	return nil
}
