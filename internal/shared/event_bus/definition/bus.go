package eventbus

import "context"

type EventBus interface {
	Publish(context.Context, []Event) error
	Subscribe(Type, Subscriber)
}
