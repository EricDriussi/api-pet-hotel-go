package service

import (
	"context"

	"github.com/EricDriussi/api-pet-hotel-go/internal/domain/booking"
)

type Booking struct {
	bookingRepository domain.BookingRepository
}

func NewBooking(bookingRepository domain.BookingRepository) Booking {
	return Booking{
		bookingRepository: bookingRepository,
	}
}

func (s Booking) RegisterBooking(ctx context.Context, id, name, duration string) error {
	booking, err := domain.NewBooking(id, name, duration)
	if err != nil {
		return err
	}
	return s.bookingRepository.Save(ctx, booking)
}
