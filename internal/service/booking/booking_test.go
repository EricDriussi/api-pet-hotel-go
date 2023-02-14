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

	t.Run("succeeds when repository returns no error", func(t *testing.T) {
		bookingRepositoryMock := new(mocks.BookingRepository)
		bookingRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Booking")).Return(nil)

		eventBusMock := new(mocks.EventBus)

		eventBusMock.On("Publish", mock.Anything, mock.MatchedBy(func(eve []eventbus.Event) bool {
			evt := eve[0]
			p := evt.(events.BookingCreatedEvent)
			return p.PetName == petName
		})).Return(nil)

		eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		bookingService := service.NewBooking(bookingRepositoryMock, eventBusMock)

		err := bookingService.RegisterBooking(context.Background(), bookingID, petName, bookingDuration)

		bookingRepositoryMock.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("fails when repository returns error", func(t *testing.T) {
		bookingRepositoryMock := new(mocks.BookingRepository)
		bookingRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Booking")).Return(errors.New("something unexpected happened"))

		eventBusMock := new(mocks.EventBus)

		bookingService := service.NewBooking(bookingRepositoryMock, eventBusMock)

		err := bookingService.RegisterBooking(context.Background(), bookingID, petName, bookingDuration)

		bookingRepositoryMock.AssertExpectations(t)
		eventBusMock.AssertExpectations(t)
		assert.Error(t, err)
	})
}
