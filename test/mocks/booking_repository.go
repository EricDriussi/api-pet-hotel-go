// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/EricDriussi/api-pet-hotel-go/internal/domain/booking"
	mock "github.com/stretchr/testify/mock"
)

// BookingRepository is an autogenerated mock type for the BookingRepository type
type BookingRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: ctx, booking
func (_m *BookingRepository) Save(ctx context.Context, booking domain.Booking) error {
	ret := _m.Called(ctx, booking)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Booking) error); ok {
		r0 = rf(ctx, booking)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBookingRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookingRepository creates a new instance of BookingRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookingRepository(t mockConstructorTestingTNewBookingRepository) *BookingRepository {
	mock := &BookingRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
