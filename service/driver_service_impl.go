package service

import (
	"TripManagementSystem/data/request"
	"TripManagementSystem/model"
	"TripManagementSystem/repository"

	validator "github.com/go-playground/validator/v10"
)

type DriverServiceImpl struct {
	DriverRepository repository.DriverRepository
	Validate         *validator.Validate
}

func NewDriverServiceImpl(driverRepository repository.DriverRepository, validate *validator.Validate) DriverService {
	return &DriverServiceImpl{
		DriverRepository: driverRepository,
		Validate:         validate,
	}
}

func (t *DriverServiceImpl) Create(driver request.CreateDriverRequest) (int64, error) {
	err := t.Validate.Struct(driver)
	if err != nil {
		return 0, err
	}

	driverModel := model.Drivers{
		FirstName:     driver.FirstName,
		LastName:      driver.LastName,
		Surname:       driver.Surname,
		LicenseNumber: driver.LicenseNumber,
		Email:         driver.Email,
		Gender:        driver.Gender,
		Dob:           driver.Dob,
		MobileNo:      driver.MobileNo,
		Address:       driver.Address,
		Language:      driver.Language,
		AccountNumber: driver.AccountNumber,
		DeviceID:      driver.DeviceID,
		DeviceType:    driver.DeviceType,
	}

	driverID, err := t.DriverRepository.Save(driverModel)
	if err != nil {
		return 0, err
	}

	return driverID, nil
}
