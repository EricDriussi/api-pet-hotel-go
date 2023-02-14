package handlers

import (
	"context"
	"errors"

	service "github.com/EricDriussi/api-pet-hotel-go/internal/service/booking"
	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/command_bus/commands"
	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/command_bus/definition"
)

type CreateBookingCommandHandler struct {
	service service.Booking
}

func NewCreateBooking(service service.Booking) CreateBookingCommandHandler {
	return CreateBookingCommandHandler{
		service: service,
	}
}

func (h CreateBookingCommandHandler) Handle(ctx context.Context, cmd commandbus.Command) error {
	createBookingCmd, ok := cmd.(commands.CreateBookingCommand)
	if !ok {
		// TODO.test?
		return errors.New("unexpected command")
	}

	return h.service.RegisterBooking(
		ctx,
		createBookingCmd.Id,
		createBookingCmd.PetName,
		createBookingCmd.Duration,
	)
}
