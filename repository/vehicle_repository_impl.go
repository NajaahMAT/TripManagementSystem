package repository

import (
	"TripManagementSystem/model"

	"gorm.io/gorm"
)

type VehicleRepositoryImpl struct {
	Db *gorm.DB
}

func NewVehicleRepositoryImpl(Db *gorm.DB) VehicleRepository {
	return &VehicleRepositoryImpl{Db: Db}
}

func (v *VehicleRepositoryImpl) Save(vehicle model.Vehicles) (int64, error) {
	tx := v.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, err
	}

	if err := tx.Create(&vehicle).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return vehicle.VehicleID, nil
}
