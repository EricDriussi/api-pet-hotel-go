package services

import (
	"context"

	"github.com/EricDriussi/api-pet-hotel-go/internal/domain/booking"
	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/event_bus"
)

type Booking struct {
	bookingRepository domain.BookingRepository
	eventBus          eventbus.EventBus
}

func NewBookingCreator(bookingRepository domain.BookingRepository, eventBus eventbus.EventBus) Booking {
	return Booking{
		bookingRepository: bookingRepository,
		eventBus:          eventBus,
	}
}

func (s Booking) RegisterBooking(ctx context.Context, id, name, duration string) error {
	booking, err := domain.NewBooking(id, name, duration)
	if err != nil {
		return err
	}
	if err := s.bookingRepository.Save(ctx, booking); err != nil {
		return err
	}
	return s.eventBus.Publish(ctx, booking.PullEvents())
}
