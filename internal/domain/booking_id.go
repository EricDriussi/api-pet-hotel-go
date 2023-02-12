package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidBookingID = errors.New("invalid Booking ID")

type BookingID struct {
	value string
}

func NewBookingID(value string) (BookingID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return BookingID{}, fmt.Errorf("%w: %s", ErrInvalidBookingID, value)
	}

	return BookingID{
		value: v.String(),
	}, nil
}

func (id BookingID) String() string {
	return id.value
}
