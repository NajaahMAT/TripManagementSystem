package service

import (
	"TripManagementSystem/data/request"
)

type TripService interface {
	Create(trip request.CreateTripRequest) (int64, error)
	UpdateTravelStatus(trips request.UpdateTripStatusRequest) error
}
