package controllers

import (
	"errors"
	"net/http"

	"github.com/EricDriussi/api-pet-hotel-go/internal/application"
	"github.com/EricDriussi/api-pet-hotel-go/internal/domain"
	"github.com/gin-gonic/gin"
)

type postBookingRequest struct {
	ID       string `json:"id" binding:"required"`
	PetName  string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func PostBooking(applicationService application.Booking) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req postBookingRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := applicationService.CreateBooking(ctx, req.ID, req.PetName, req.Duration)
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
