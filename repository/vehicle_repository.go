package repository

import "TripManagementSystem/model"

type VehicleRepository interface {
	Save(vehicles model.Vehicles) (int64, error)
}
