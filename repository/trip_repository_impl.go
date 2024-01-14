package repository

import (
	"TripManagementSystem/model"

	"gorm.io/gorm"
)

type TripRepositoryImpl struct {
	Db *gorm.DB
}

func NewTripRepositoryImpl(Db *gorm.DB) TripRepository {
	return &TripRepositoryImpl{Db: Db}
}

func (t *TripRepositoryImpl) Save(trip model.Trips) (int64, error) {
	tx := t.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return 0, err
	}

	if err := tx.Create(&trip).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return trip.TripID, nil
}

func (t *TripRepositoryImpl) UpdateTravelStatus(trip model.Trips) error {
	tx := t.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&model.Trips{}).Where("trip_id = ?", trip.TripID).Update("trip_status", trip.TripStatus).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
