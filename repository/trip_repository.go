package repository

import "TripManagementSystem/model"

type TripRepository interface {
	Save(trips model.Trips) (int64, error)
	UpdateTravelStatus(trips model.Trips) error
	// FindById(tripsId int64) (trips model.Trips, err error)
}
