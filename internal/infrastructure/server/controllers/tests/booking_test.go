package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/server/controllers"
	"github.com/EricDriussi/api-pet-hotel-go/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_PostBooking(t *testing.T) {
	commandBus := new(mocks.CommandBus)
	commandBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("commands.CreateBookingCommand"),
	).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/booking", controllers.PostBooking(commandBus))

	t.Run("return 201 when given a valid request", func(t *testing.T) {
		createBookingReq := controllers.PostBookingRequest{
			ID:       "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			PetName:  "A Pet",
			Duration: "1 months",
		}

		b, err := json.Marshal(createBookingReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/booking", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("returns 400 when given a partial request", func(t *testing.T) {
		createBookingReq := controllers.PostBookingRequest{
			PetName:  "A Pet",
			Duration: "1 months",
		}

		b, err := json.Marshal(createBookingReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/booking", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
