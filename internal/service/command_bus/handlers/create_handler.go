package handlers

import (
	"context"
	"errors"

	service "github.com/EricDriussi/api-pet-hotel-go/internal/service/booking"
	command "github.com/EricDriussi/api-pet-hotel-go/internal/service/command_bus"
	"github.com/EricDriussi/api-pet-hotel-go/internal/service/command_bus/commands"
)

type CreateBookingCommandHandler struct {
	service service.Booking
}

func NewCreateBooking(service service.Booking) CreateBookingCommandHandler {
	return CreateBookingCommandHandler{
		service: service,
	}
}

func (h CreateBookingCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
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
