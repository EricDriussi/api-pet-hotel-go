package handlers

import (
	"context"
	"errors"

	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/commands"
	"github.com/EricDriussi/api-pet-hotel-go/internal/services"
	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/command_bus"
)

type CreateBookingCommandHandler struct {
	service services.Booking
}

func NewCreateBooking(service services.Booking) CreateBookingCommandHandler {
	return CreateBookingCommandHandler{
		service: service,
	}
}

func (h CreateBookingCommandHandler) Handle(ctx context.Context, cmd commandbus.Command) error {
	createBookingCmd, ok := cmd.(commands.CreateBookingCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.RegisterBooking(
		ctx,
		createBookingCmd.Id,
		createBookingCmd.PetName,
		createBookingCmd.Duration,
	)
}
