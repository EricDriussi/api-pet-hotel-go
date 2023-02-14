package controllers

import (
	"errors"
	"net/http"

	domain "github.com/EricDriussi/api-pet-hotel-go/internal/domain/booking"
	"github.com/EricDriussi/api-pet-hotel-go/internal/shared/command_bus/commands"
	commandbus "github.com/EricDriussi/api-pet-hotel-go/internal/shared/command_bus/definition"
	"github.com/gin-gonic/gin"
)

type PostBookingRequest struct {
	ID       string `json:"id" binding:"required"`
	PetName  string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func PostBooking(commandBus commandbus.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req PostBookingRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := commandBus.Dispatch(ctx, commands.NewCreateBooking(
			req.ID,
			req.PetName,
			req.Duration,
		))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidBookingID),
				errors.Is(err, domain.ErrEmptyPetName), errors.Is(err, domain.ErrInvalidBookingID):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
