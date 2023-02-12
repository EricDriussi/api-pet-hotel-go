package domain

import "errors"

var ErrEmptyDuration = errors.New("the field Duration can not be empty")

type BookingDuration struct {
	value string
}

func NewBookingDuration(value string) (BookingDuration, error) {
	if value == "" {
		return BookingDuration{}, ErrEmptyDuration
	}

	return BookingDuration{
		value: value,
	}, nil
}

func (duration BookingDuration) String() string {
	return duration.value
}
