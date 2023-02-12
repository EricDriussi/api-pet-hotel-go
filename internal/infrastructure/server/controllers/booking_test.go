package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EricDriussi/api-pet-hotel-go/internal/application"
	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/server/controllers"
	"github.com/EricDriussi/api-pet-hotel-go/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_PostBooking(t *testing.T) {
	repositoryMock := new(mocks.BookingRepository)
	repositoryMock.On(
		"Save",
		mock.Anything,
		mock.Anything,
	).Return(nil)

	bookingService := application.NewBookingService(repositoryMock)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/booking", controllers.PostBooking(bookingService))

	t.Run("given a partial request it returns 400", func(t *testing.T) {
		createCourseReq := controllers.PostBookingRequest{
			PetName:  "A Pet",
			Duration: "1 months",
		}

		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/booking", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		createCourseReq := controllers.PostBookingRequest{
			ID:       "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			PetName:  "A Pet",
			Duration: "1 months",
		}

		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/booking", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("given request with invalid id returns 400", func(t *testing.T) {
		createCourseReq := controllers.PostBookingRequest{
			ID:       "ba57",
			PetName:  "A Pet",
			Duration: "1 months",
		}

		b, err := json.Marshal(createCourseReq)
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
