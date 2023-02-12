package server

import (
	"fmt"
	"log"

	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/server/controllers"
	commandbus "github.com/EricDriussi/api-pet-hotel-go/internal/service/command_bus"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	commandBus commandbus.Bus
}

func New(host string, port uint, bookingService commandbus.Bus) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		commandBus: bookingService,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", controllers.Health)
	s.engine.POST("/booking", controllers.PostBooking(s.commandBus))
}
