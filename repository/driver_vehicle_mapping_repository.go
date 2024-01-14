package repository

import "TripManagementSystem/model"

type DVMappingRepository interface {
	Save(vehicles model.DriverVehicleMappings) (int64, error)
}
