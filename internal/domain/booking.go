package domain

import (
	"context"
)

type Booking struct {
	ID       BookingID
	PetName  PetName
	Duration BookingDuration
}

type BookingRepository interface {
	Save(ctx context.Context, booking Booking) error
}

func NewBooking(id, name, duration string) (Booking, error) {
	idVO, err := NewBookingID(id)
	if err != nil {
		return Booking{}, err
	}
	nameVO, err := NewPetName(name)
	if err != nil {
		return Booking{}, err
	}
	durationVO, err := NewBookingDuration(duration)
	if err != nil {
		return Booking{}, err
	}

	return Booking{
		ID:       idVO,
		PetName:  nameVO,
		Duration: durationVO,
	}, nil
}
