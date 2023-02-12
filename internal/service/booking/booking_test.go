package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/EricDriussi/api-pet-hotel-go/internal/domain/booking"
	"github.com/EricDriussi/api-pet-hotel-go/internal/service/booking"
	"github.com/EricDriussi/api-pet-hotel-go/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_BookingService_CreateCourse_Succeed(t *testing.T) {
	bookingID, petName, bookingDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Nice Pet Name", "1 months"

	booking, err := domain.NewBooking(bookingID, petName, bookingDuration)
	require.NoError(t, err)

	bookingRepositoryMock := new(mocks.BookingRepository)
	bookingRepositoryMock.On("Save", mock.Anything, booking).Return(nil)

	bookingService := service.NewBooking(bookingRepositoryMock)

	err = bookingService.RegisterBooking(context.Background(), bookingID, petName, bookingDuration)

	bookingRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_BookingService_CreateCourse_RepositoryError(t *testing.T) {
	bookingID, petName, bookingDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Nice Pet Name", "1 months"
	booking, err := domain.NewBooking(bookingID, petName, bookingDuration)
	require.NoError(t, err)

	bookingRepositoryMock := new(mocks.BookingRepository)
	bookingRepositoryMock.On("Save", mock.Anything, booking).Return(errors.New("something unexpected happened"))

	bookingService := service.NewBooking(bookingRepositoryMock)

	err = bookingService.RegisterBooking(context.Background(), bookingID, petName, bookingDuration)

	bookingRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}
