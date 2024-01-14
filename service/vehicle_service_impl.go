package service

import (
	"TripManagementSystem/data/request"
	"TripManagementSystem/model"
	"TripManagementSystem/repository"

	validator "github.com/go-playground/validator/v10"
)

type VehicleServiceImpl struct {
	VehicleRepository repository.VehicleRepository
	Validate          *validator.Validate
}

func NewVehicleServiceImpl(vehicleRepository repository.VehicleRepository, validate *validator.Validate) VehicleService {
	return &VehicleServiceImpl{
		VehicleRepository: vehicleRepository,
		Validate:          validate,
	}
}

func (t *VehicleServiceImpl) Create(vehicle request.CreateVehicleRequest) (int64, error) {
	err := t.Validate.Struct(vehicle)
	if err != nil {
		return 0, err
	}

	vehicleModel := model.Vehicles{
		RegistrationNumber: vehicle.RegistrationNumber,
		Type:               vehicle.Type,
		ModelMake:          vehicle.ModelMake,
		BrandModel:         vehicle.BrandModel,
		Color:              vehicle.Color,
		Year:               vehicle.Year,
		SeatingCapacity:    vehicle.SeatingCapacity,
		EngineCapacity:     vehicle.EngineCapacity,
	}

	vehicleID, err := t.VehicleRepository.Save(vehicleModel)
	if err != nil {
		return 0, err
	}

	return vehicleID, nil
}
