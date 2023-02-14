package inmemory

import (
	"context"
	"log"

	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/command_bus/definition"
)

type CommandBus struct {
	handlers map[commandbus.Type]commandbus.Handler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[commandbus.Type]commandbus.Handler),
	}
}

func (b *CommandBus) Dispatch(ctx context.Context, cmd commandbus.Command) error {
	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		log.Printf("No handler for %s", cmd.Type())
		// TODO. return err and test
		return nil
	}

	if err := handler.Handle(ctx, cmd); err != nil {
		log.Printf("Error while handling %s - %s\n", cmd.Type(), err)
		return err
	}

	return nil
}

func (b *CommandBus) Register(cmdType commandbus.Type, handler commandbus.Handler) {
	b.handlers[cmdType] = handler
}
