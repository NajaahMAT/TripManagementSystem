package service

import (
	"TripManagementSystem/data/request"
)

type DriverService interface {
	Create(driver request.CreateDriverRequest) (int64, error)
}
