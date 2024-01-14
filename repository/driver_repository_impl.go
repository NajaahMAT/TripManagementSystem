package repository

import (
	"TripManagementSystem/model"

	"gorm.io/gorm"
)

type DriverRepositoryImpl struct {
	Db *gorm.DB
}

func NewDriverRepositoryImpl(Db *gorm.DB) DriverRepository {
	return &DriverRepositoryImpl{Db: Db}
}

func (d *DriverRepositoryImpl) Save(driver model.Drivers) (int64, error) {
	tx := d.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, err
	}

	if err := tx.Create(&driver).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return driver.DriverID, nil
}
