package mysqlrepo_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/EricDriussi/api-pet-hotel-go/internal/domain/booking"
	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/repositories/mysql_repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_BookingRepository_Save(t *testing.T) {
	bookingID, petName, bookingDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Nice Pet Name", "1 months"

	booking, err := domain.NewBooking(bookingID, petName, bookingDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	t.Run("returns no error if succeeds", func(t *testing.T) {
		sqlMock.ExpectExec(
			"INSERT INTO bookings (id, pet_name, duration) VALUES (?, ?, ?)").
			WithArgs(bookingID, petName, bookingDuration).
			WillReturnResult(sqlmock.NewResult(0, 1))

		repo := mysqlrepo.NewBooking(db)

		err = repo.Save(context.Background(), booking)

		assert.NoError(t, sqlMock.ExpectationsWereMet())
		assert.NoError(t, err)
	})

	t.Run("returns an error if fails", func(t *testing.T) {
		sqlMock.ExpectExec(
			"INSERT INTO bookings (id, pet_name, duration) VALUES (?, ?, ?)").
			WithArgs(bookingID, petName, bookingDuration).
			WillReturnError(errors.New("FAILED"))

		repo := mysqlrepo.NewBooking(db)

		err = repo.Save(context.Background(), booking)

		assert.NoError(t, sqlMock.ExpectationsWereMet())
		assert.Error(t, err)
	})
}
