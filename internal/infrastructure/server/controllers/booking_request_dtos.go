package controllers

type PostBookingRequest struct {
	ID       string `json:"id" binding:"required"`
	PetName  string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}
