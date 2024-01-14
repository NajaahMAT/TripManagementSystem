package service

import (
	"TripManagementSystem/data/request"
)

type VehicleService interface {
	Create(vehicle request.CreateVehicleRequest) (int64, error)
}
