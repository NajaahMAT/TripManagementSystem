package repository

import (
	"TripManagementSystem/model"
	"errors"

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

	// Fetch the current state of the trip to check its current status
	var currentTrip model.Trips
	if err := tx.Where("trip_id = ?", trip.TripID).First(&currentTrip).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Check if the trip is in a state that allows status update
	if (currentTrip.TripStatus != "Created" && trip.TripStatus == "InProgress") ||
		(currentTrip.TripStatus != "InProgress" && trip.TripStatus == "Completed") {
		tx.Rollback()
		return errors.New("trip status cannot be updated from its current state")
	}

	// Proceed with updating the status
	if err := tx.Model(&model.Trips{}).Where("trip_id = ?", trip.TripID).Update("trip_status", trip.TripStatus).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
