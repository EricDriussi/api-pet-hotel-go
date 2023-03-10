package events

import "github.com/EricDriussi/api-pet-hotel-go/internal/shared/event_bus"

const BookingCreatedEventType eventbus.Type = "events.booking.created"

type BookingCreatedEvent struct {
	eventbus.BaseEvent
	Id       string
	PetName  string
	Duration string
}

func NewBookingCreated(id, name, duration string) BookingCreatedEvent {
	return BookingCreatedEvent{
		Id:       id,
		PetName:  name,
		Duration: duration,

		BaseEvent: eventbus.NewBaseEvent(id),
	}
}

func (e BookingCreatedEvent) Type() eventbus.Type {
	return BookingCreatedEventType
}
