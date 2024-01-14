package service

import (
	"TripManagementSystem/data/request"
)

type DMMappingService interface {
	Create(vehicle request.CreateDVMappingRequest) (int64, error)
}
