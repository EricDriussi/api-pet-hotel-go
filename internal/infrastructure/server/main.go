package server

import (
	"fmt"
	"log"

	"github.com/EricDriussi/api-pet-hotel-go/internal/application"
	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/server/controllers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	applicationService application.Booking
}

func New(host string, port uint, bookingService application.Booking) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		applicationService: bookingService,
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
	s.engine.POST("/booking", controllers.PostBooking(s.applicationService))
}
