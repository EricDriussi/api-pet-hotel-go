package domain

type bookingDuration struct {
	value string
}

func newDuration(value string) (bookingDuration, error) {
	if value == "" {
		return bookingDuration{}, ErrEmptyDuration
	}

	return bookingDuration{
		value: value,
	}, nil
}

func (duration bookingDuration) string() string {
	return duration.value
}
