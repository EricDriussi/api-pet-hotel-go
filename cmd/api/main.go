package api

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/command_bus/in_memory"
	mysqlrepo "github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/repositories/mysql_repo"
	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/server"
	service "github.com/EricDriussi/api-pet-hotel-go/internal/service/booking"
	"github.com/EricDriussi/api-pet-hotel-go/internal/service/command_bus/commands"
	"github.com/EricDriussi/api-pet-hotel-go/internal/service/command_bus/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func Deploy() {
	if err := bootstrap(); err != nil {
		log.Fatal(err)
	}
}

const (
	host   = "localhost"
	port   = 6969
	dbUser = "root"
	dbPass = "root"
	dbHost = "localhost"
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
	bookingService := service.NewBooking(bookingRepository)

	inMemoryCommandBus := inmemory.NewCommandBus()
	createBookingCommandHandler := handlers.NewCreateBooking(bookingService)
	inMemoryCommandBus.Register(commands.CreateBookingCommandType, createBookingCommandHandler)

	srv := server.New(host, port, inMemoryCommandBus)
	return srv.Run()
}
