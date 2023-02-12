package domain

import (
	"context"
	"errors"
)

type BookingRepository interface {
	Save(ctx context.Context, booking Booking) error
}

var (
	ErrInvalidBookingID = errors.New("invalid Booking ID")
	ErrEmptyPetName     = errors.New("Pet Name can not be empty")
	ErrEmptyDuration    = errors.New("Duration can not be empty")
)

type Booking struct {
	id       bookingID
	petName  petName
	duration bookingDuration
}

func NewBooking(id, name, duration string) (Booking, error) {
	idVO, err := newID(id)
	if err != nil {
		return Booking{}, err
	}
	nameVO, err := newPetName(name)
	if err != nil {
		return Booking{}, err
	}
	durationVO, err := newDuration(duration)
	if err != nil {
		return Booking{}, err
	}

	return Booking{
		id:       idVO,
		petName:  nameVO,
		duration: durationVO,
	}, nil
}

func (b Booking) IDAsString() string {
	return b.id.string()
}

func (b Booking) DurationAsString() string {
	return b.duration.string()
}

func (b Booking) PetNameAsString() string {
	return b.petName.string()
}
