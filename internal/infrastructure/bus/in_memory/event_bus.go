package inmemory

import (
	"context"

	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/event_bus"
)

type EventBus struct {
	handlers map[eventbus.Type][]eventbus.Subscriber
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[eventbus.Type][]eventbus.Subscriber),
	}
}

func (b *EventBus) Publish(ctx context.Context, events []eventbus.Event) error {
	for _, evt := range events {
		handlers, ok := b.handlers[evt.Type()]
		if !ok {
			return nil
		}

		for _, handler := range handlers {
			handler.Handle(ctx, evt)
		}
	}

	return nil
}

func (b *EventBus) Subscribe(evtType eventbus.Type, handler eventbus.Subscriber) {
	subscribersForType, ok := b.handlers[evtType]
	if !ok {
		b.handlers[evtType] = []eventbus.Subscriber{handler}
	}

	subscribersForType = append(subscribersForType, handler)
}
