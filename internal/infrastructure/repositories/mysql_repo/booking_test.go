package repositories_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/EricDriussi/api-pet-hotel-go/internal/domain"
	repositories "github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/repositories/mysql_repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_BookingRepository_Save_Succeed(t *testing.T) {
	bookingID, petName, bookingDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Nice Pet Name", "1 months"

	booking, err := domain.NewBooking(bookingID, petName, bookingDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO bookings (id, pet_name, duration) VALUES (?, ?, ?)").
		WithArgs(bookingID, petName, bookingDuration).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repositories.NewBookingRepo(db)

	err = repo.Save(context.Background(), booking)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func Test_BookingRepository_Save_RepositoryError(t *testing.T) {
	bookingID, petName, bookingDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Nice Pet Name", "1 months"

	booking, err := domain.NewBooking(bookingID, petName, bookingDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO bookings (id, pet_name, duration) VALUES (?, ?, ?)").
		WithArgs(bookingID, petName, bookingDuration).
		WillReturnError(errors.New("FAILED"))

	repo := repositories.NewBookingRepo(db)

	err = repo.Save(context.Background(), booking)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}
