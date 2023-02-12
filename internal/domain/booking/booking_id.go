package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type bookingID struct {
	value string
}

func newID(value string) (bookingID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return bookingID{}, fmt.Errorf("%w: %s", ErrInvalidBookingID, value)
	}

	return bookingID{
		value: v.String(),
	}, nil
}

func (id bookingID) string() string {
	return id.value
}
