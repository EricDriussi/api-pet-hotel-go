package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	mysqlrepo "github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/repositories/mysql_repo"
	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/server"
	service "github.com/EricDriussi/api-pet-hotel-go/internal/service/booking"
	counterService "github.com/EricDriussi/api-pet-hotel-go/internal/service/counter"
	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/command_bus/commands"
	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/command_bus/handlers"
	inmemory_commandbus "github.com/EricDriussi/api-pet-hotel-go/internal/shared/command_bus/in_memory"
	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/event_bus/events"
	inmemory_eventbus "github.com/EricDriussi/api-pet-hotel-go/internal/shared/event_bus/in_memory"
	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/event_bus/subscribers"

	_ "github.com/go-sql-driver/mysql"
)

func Deploy() {
	if err := bootstrap(); err != nil {
		log.Fatal(err)
	}
}

const (
	host   = "0.0.0.0"
	port   = 6969
	dbUser = "admin"
	dbPass = "admin"
	dbHost = "mysql"
	dbPort = "3306"
	dbName = "pet_hotel"
)

func bootstrap() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	bookingRepository := mysqlrepo.NewBooking(db)
	inMemoryEventBus := inmemory_eventbus.NewEventBus()
	bookingService := service.NewBooking(bookingRepository, inMemoryEventBus)

	counterService := counterService.NewBookingCounter()
	inMemoryEventBus.Subscribe(events.BookingCreatedEventType, subscribers.NewBookingCreatedSubscriber(counterService))

	inMemoryCommandBus := inmemory_commandbus.NewCommandBus()
	createBookingCommandHandler := handlers.NewCreateBooking(bookingService)
	inMemoryCommandBus.Register(commands.CreateBookingCommandType, createBookingCommandHandler)

	// srv := server.New(host, port, inMemoryCommandBus)
	// return srv.Run()

	cleanContext := context.Background()
	srv := server.NewGraceful(cleanContext, host, port, inMemoryCommandBus)
	return srv.Run()
}
