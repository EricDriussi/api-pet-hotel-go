package events

import eventbus "github.com/EricDriussi/api-pet-hotel-go/internal/shared/event_bus/definition"

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
