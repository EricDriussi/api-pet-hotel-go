package domain

import (
	"context"
	"errors"
)

var ErrInvalidBookingID = errors.New("invalid Booking ID")

var ErrEmptyPetName = errors.New("the field Pet Name can not be empty")

var ErrEmptyDuration = errors.New("the field Duration can not be empty")

// TODO.VO
type Booking struct {
	ID       string
	PetName  string
	Duration string
}

type BookingRepository interface {
	Save(ctx context.Context, booking Booking) error
}

func NewBooking(id, name, duration string) (Booking, error) {
	return Booking{
		ID:       id,
		PetName:  name,
		Duration: duration,
	}, nil
}
