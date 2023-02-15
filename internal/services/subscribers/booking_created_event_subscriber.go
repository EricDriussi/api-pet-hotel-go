package subscribers

import (
	"context"
	"errors"

	"github.com/EricDriussi/api-pet-hotel-go/internal/domain/events"
	"github.com/EricDriussi/api-pet-hotel-go/internal/services"
	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/event_bus"
)

type BookingCreatedSubscriber struct {
	service services.BookingCounterService
}

func NewBookingCreatedSubscriber(service services.BookingCounterService) BookingCreatedSubscriber {
	return BookingCreatedSubscriber{service: service}
}

func (s BookingCreatedSubscriber) Handle(_ context.Context, event eventbus.Event) error {
	courseCreatedEvt, ok := event.(events.BookingCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return s.service.Increase(courseCreatedEvt.ID())
}
