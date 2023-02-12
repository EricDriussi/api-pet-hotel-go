package controllers

import (
	"errors"
	"net/http"

	"github.com/EricDriussi/api-pet-hotel-go/internal/domain/booking"
	"github.com/EricDriussi/api-pet-hotel-go/internal/service/booking"
	"github.com/gin-gonic/gin"
)

type PostBookingRequest struct {
	ID       string `json:"id" binding:"required"`
	PetName  string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func PostBooking(applicationService service.Booking) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req PostBookingRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := applicationService.RegisterBooking(ctx, req.ID, req.PetName, req.Duration)
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
