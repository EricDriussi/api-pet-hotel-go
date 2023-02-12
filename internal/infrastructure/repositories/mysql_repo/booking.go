package mysqlrepo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/EricDriussi/api-pet-hotel-go/internal/domain/booking"
	"github.com/huandu/go-sqlbuilder"
)

const (
	sqlTable = "bookings"
)

type sqlBooking struct {
	ID       string `db:"id"`
	PetName  string `db:"pet_name"`
	Duration string `db:"duration"`
}

type bookingRepo struct {
	db *sql.DB
}

func NewBooking(db *sql.DB) *bookingRepo {
	return &bookingRepo{db: db}
}

func (r *bookingRepo) Save(ctx context.Context, booking domain.Booking) error {
	bookingSQLStruct := sqlbuilder.NewStruct(new(sqlBooking))
	query, args := bookingSQLStruct.InsertInto(sqlTable, sqlBooking{
		ID:       booking.IDAsString(),
		PetName:  booking.PetNameAsString(),
		Duration: booking.DurationAsString(),
	}).Build()

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("error persisting booking: %v", err)
	}

	return nil
}
