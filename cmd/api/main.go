package api

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/repositories/mysql_repo"
	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/server"
	"github.com/EricDriussi/api-pet-hotel-go/internal/service/booking"

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

	srv := server.New(host, port, bookingService)
	return srv.Run()
}
