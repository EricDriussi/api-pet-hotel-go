package application_test

import (
	"context"
	"errors"
	"testing"

	"github.com/EricDriussi/api-pet-hotel-go/internal/application"
	"github.com/EricDriussi/api-pet-hotel-go/internal/domain"
	"github.com/EricDriussi/api-pet-hotel-go/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// mockery -r --case=snake --outpkg=mocks --output=test/mocks --name=BookingRepository
//go:generate mockery --case=snake --outpkg=mocks --output=test/mocks --name=BookingRepo
//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

func Test_BookingService_CreateCourse_Succeed(t *testing.T) {
	bookingID, petName, bookingDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Nice Pet Name", "1 months"

	booking, err := domain.NewBooking(bookingID, petName, bookingDuration)
	require.NoError(t, err)

	bookingRepositoryMock := new(mocks.BookingRepository)
	bookingRepositoryMock.On("Save", mock.Anything, booking).Return(nil)

	bookingService := application.NewBookingService(bookingRepositoryMock)

	err = bookingService.CreateBooking(context.Background(), bookingID, petName, bookingDuration)

	bookingRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_BookingService_CreateCourse_RepositoryError(t *testing.T) {
	bookingID, petName, bookingDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Nice Pet Name", "1 months"
	booking, err := domain.NewBooking(bookingID, petName, bookingDuration)
	require.NoError(t, err)

	bookingRepositoryMock := new(mocks.BookingRepository)
	bookingRepositoryMock.On("Save", mock.Anything, booking).Return(errors.New("something unexpected happened"))

	bookingService := application.NewBookingService(bookingRepositoryMock)

	err = bookingService.CreateBooking(context.Background(), bookingID, petName, bookingDuration)

	bookingRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}
