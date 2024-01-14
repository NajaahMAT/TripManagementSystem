package repository

import (
	"TripManagementSystem/model"

	"gorm.io/gorm"
)

type DVMappingRepositoryImpl struct {
	Db *gorm.DB
}

func NewDVMappingRepositoryImpl(Db *gorm.DB) DVMappingRepository {
	return &DVMappingRepositoryImpl{Db: Db}
}

func (v *DVMappingRepositoryImpl) Save(dvMapping model.DriverVehicleMappings) (int64, error) {
	tx := v.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, err
	}

	if err := tx.Create(&dvMapping).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return dvMapping.MappingID, nil
}
