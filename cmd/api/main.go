package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/EricDriussi/api-pet-hotel-go/internal/domain/events"
	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/bus/in_memory"
	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/commands"
	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/repositories/mysql_repo"
	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/server"
	"github.com/EricDriussi/api-pet-hotel-go/internal/services"
	"github.com/EricDriussi/api-pet-hotel-go/internal/services/handlers"
	"github.com/EricDriussi/api-pet-hotel-go/internal/services/subscribers"

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
	inMemoryEventBus := inmemory.NewEventBus()
	bookingService := services.NewBookingCreator(bookingRepository, inMemoryEventBus)

	counterService := services.NewBookingCounter()
	inMemoryEventBus.Subscribe(events.BookingCreatedEventType, subscribers.NewBookingCreatedSubscriber(counterService))

	inMemoryCommandBus := inmemory.NewCommandBus()
	inMemoryCommandBus.Register(commands.CreateBookingCommandType, handlers.NewCreateBooking(bookingService))

	// srv := server.New(host, port, inMemoryCommandBus)
	// return srv.Run()

	cleanContext := context.Background()
	srv := server.NewGraceful(cleanContext, host, port, inMemoryCommandBus)
	return srv.Run()
}
