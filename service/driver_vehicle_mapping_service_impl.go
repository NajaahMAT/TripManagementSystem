package service

import (
	"TripManagementSystem/data/request"
	"TripManagementSystem/model"
	"TripManagementSystem/repository"

	validator "github.com/go-playground/validator/v10"
)

type DVMappingServiceImpl struct {
	DVMappingRepository repository.DVMappingRepository
	Validate            *validator.Validate
}

func NewDVMappingServiceImpl(dvMappingRepository repository.DVMappingRepository, validate *validator.Validate) DMMappingService {
	return &DVMappingServiceImpl{
		DVMappingRepository: dvMappingRepository,
		Validate:            validate,
	}
}

func (t *DVMappingServiceImpl) Create(dvMapping request.CreateDVMappingRequest) (int64, error) {
	err := t.Validate.Struct(dvMapping)
	if err != nil {
		return 0, err
	}

	dvMappingModel := model.DriverVehicleMappings{
		DriverID:  dvMapping.DriverID,
		VehicleID: dvMapping.VehicleID,
		StartedAt: dvMapping.MappingStartedAt,
		EndedAt:   dvMapping.MappingEndedAt,
	}

	mappingID, err := t.DVMappingRepository.Save(dvMappingModel)
	if err != nil {
		return 0, err // Return error if save operation failed
	}

	return mappingID, nil
}
