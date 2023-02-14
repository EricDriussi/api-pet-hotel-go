package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/server/controllers"
	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/server/middleware"
	commandbus "github.com/EricDriussi/api-pet-hotel-go/internal/shared/command_bus/definition"
	"github.com/gin-gonic/gin"
)

type GracefulServer struct {
	httpAddr        string
	engine          *gin.Engine
	context         context.Context
	shutdownTimeout time.Duration
	commandBus      commandbus.CommandBus
}

func (s *GracefulServer) registerRoutes() {
	s.engine.GET("/health", controllers.Health)
	s.engine.POST("/booking", controllers.PostBooking(s.commandBus))
}

func NewGraceful(ctx context.Context, host string, port uint, commandBus commandbus.CommandBus) GracefulServer {
	engine_with_middlewares := gin.New()
	engine_with_middlewares.Use(gin.Recovery(), gin.Logger()) // same as gin.Default()
	engine_with_middlewares.Use(middleware.DiscoInferno)
	srv := GracefulServer{
		engine:   engine_with_middlewares,
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		context:  cancelOnSIGINT(ctx),

		shutdownTimeout: 10 * time.Second,

		commandBus: commandBus,
	}

	srv.registerRoutes()
	return srv
}

func (s *GracefulServer) Run() error {
	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-s.context.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func cancelOnSIGINT(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
