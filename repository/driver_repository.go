package repository

import "TripManagementSystem/model"

type DriverRepository interface {
	Save(drivers model.Drivers) (int64, error)
}
