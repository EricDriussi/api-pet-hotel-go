package service

type BookingCounterService struct{}

func NewBookingCounter() BookingCounterService {
	return BookingCounterService{}
}

func (s BookingCounterService) Increase(id string) error {
	return nil
}
