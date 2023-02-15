package service_test

import (
	"context"
	"errors"
	"testing"

	service "github.com/EricDriussi/api-pet-hotel-go/internal/service/booking"
	eventbus "github.com/EricDriussi/api-pet-hotel-go/internal/shared/event_bus/definition"
	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/event_bus/events"
	"github.com/EricDriussi/api-pet-hotel-go/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_BookingService_CreateBooking(t *testing.T) {
	bookingID, petName, bookingDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Nice Pet Name", "1 months"
	irrelevantMock := mock.Anything
	aBookingDomainMock := mock.AnythingOfType("domain.Booking")
	anError := errors.New("something unexpected happened")

	t.Run("succeeds when repository and event bus return no error", func(t *testing.T) {
		bookingRepositoryMock := new(mocks.BookingRepository)
		bookingRepositoryMock.On("Save", irrelevantMock, aBookingDomainMock).Return(nil)

		eventBusMock := new(mocks.EventBus)
		eventBusMock.On("Publish", irrelevantMock, mock.MatchedBy(func(eve []eventbus.Event) bool {
			evt := eve[0]
			p := evt.(events.BookingCreatedEvent)
			return p.PetName == petName
		})).Return(nil)

		bookingService := service.NewBooking(bookingRepositoryMock, eventBusMock)

		err := bookingService.RegisterBooking(context.Background(), bookingID, petName, bookingDuration)

		bookingRepositoryMock.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("fails when repository returns error", func(t *testing.T) {
		bookingRepositoryMock := new(mocks.BookingRepository)
		bookingRepositoryMock.On("Save", irrelevantMock, aBookingDomainMock).Return(anError)

		eventBusMock := new(mocks.EventBus)

		bookingService := service.NewBooking(bookingRepositoryMock, eventBusMock)

		err := bookingService.RegisterBooking(context.Background(), bookingID, petName, bookingDuration)

		bookingRepositoryMock.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("fails when event bus returns error", func(t *testing.T) {
		bookingRepositoryMock := new(mocks.BookingRepository)
		bookingRepositoryMock.On("Save", irrelevantMock, aBookingDomainMock).Return(nil)

		eventBusMock := new(mocks.EventBus)
		anEvent := mock.AnythingOfType("[]eventbus.Event")
		eventBusMock.On("Publish", irrelevantMock, anEvent).Return(anError)

		bookingService := service.NewBooking(bookingRepositoryMock, eventBusMock)

		err := bookingService.RegisterBooking(context.Background(), bookingID, petName, bookingDuration)

		eventBusMock.AssertExpectations(t)
		assert.Error(t, err)
	})
}
