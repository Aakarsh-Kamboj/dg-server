package main

import (
	"context"
	"dg-server/config"
	"dg-server/infrastructure"
	custommiddleware "dg-server/internal/middleware"
	transportHttp "dg-server/internal/transport/http"
	v1 "dg-server/internal/transport/http/v1"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// Server wires together Echo and your service.
type Server struct {
	e                   *echo.Echo
	cfg                 *config.Config
	logger              *zap.Logger
	onboardingHandler   *v1.OnboardingHandler
	controlHandler      *v1.ControlHandler
	evidenceTaskHandler *v1.EvidenceTaskHandler
	frameworkHandler    *v1.FrameworkHandler
	departmentHandler   *v1.DepartmentHandler
	// usecase *usecase.Usecase
}

// NewEcho is a provider for *echo.Echo.
func NewEcho(logger *zap.Logger) *echo.Echo {
	e := echo.New()
	// you can register global middleware here

	// Recover middleware to recover from panics anywhere in the chain
	e.Use(middleware.Recover())

	// Request ID middleware to generate a unique ID for each request
	e.Use(custommiddleware.RequestIDMiddleware(logger))

	// Request logger middleware to log incoming requests
	e.Use(custommiddleware.RequestLoggerMiddleware(logger))

	// CORS middleware to allow cross-origin requests
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// Register the validator
	e.Validator = &infrastructure.CustomValidator{Validator: validator.New()}
	return e
}

// NewServer initializes a new server instance.
// NewServer is a provider for *Server.
func NewServer(e *echo.Echo, cfg *config.Config, logger *zap.Logger, onboardingHandler *v1.OnboardingHandler, controlHandler *v1.ControlHandler, evidenceTaskHandler *v1.EvidenceTaskHandler, frameworkHandler *v1.FrameworkHandler, departmentHandler *v1.DepartmentHandler) *Server {
	server := &Server{e: e, cfg: cfg, logger: logger, onboardingHandler: onboardingHandler, controlHandler: controlHandler, evidenceTaskHandler: evidenceTaskHandler, frameworkHandler: frameworkHandler, departmentHandler: departmentHandler}
	transportHttp.RegisterRoutes(server.e, onboardingHandler, controlHandler, evidenceTaskHandler, frameworkHandler, departmentHandler)
	return server
}

// Start starts the server on the configured port.
func (s *Server) Start() error {
	port := s.cfg.Server.Port

	// Run the server in a goroutine to avoid blocking the main thread and to allow graceful shutdown
	go func() {
		// Log the server start
		s.logger.Info("Starting API server", zap.String("port", port))
		if err := s.e.Start(":" + port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatal("Unexpected error while running server", zap.Error(err))
		}
	}()

	// Wait for a signal to shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.logger.Info("Received shutdown signal, shutting down server")
	if err := s.Shutdown(s.cfg.Server.ShutdownTimeout); err != nil {
		s.logger.Fatal("Error during graceful shutdown", zap.Error(err))
	}
	return nil
}

func (s *Server) Shutdown(shutdownTimeout time.Duration) error {
	// Shutdown the server gracefully with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
	defer cancel()
	if err := s.e.Shutdown(ctx); err != nil {
		s.logger.Error("Failed to shutdown server", zap.Error(err))
		return err
	}
	s.logger.Info("Server shut down gracefully")
	return nil
}
