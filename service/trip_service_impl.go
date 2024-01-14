package service

import (
	"TripManagementSystem/data/request"
	"TripManagementSystem/model"
	"TripManagementSystem/repository"

	validator "github.com/go-playground/validator/v10"
)

type TripServiceImpl struct {
	TripRepository repository.TripRepository
	Validate       *validator.Validate
}

func NewTripServiceImpl(tripRepository repository.TripRepository, validate *validator.Validate) TripService {
	return &TripServiceImpl{
		TripRepository: tripRepository,
		Validate:       validate,
	}
}

func (t *TripServiceImpl) Create(trip request.CreateTripRequest) (int64, error) {
	err := t.Validate.Struct(trip)
	if err != nil {
		return 0, err
	}

	tripModel := model.Trips{
		VehicleID:         trip.VehicleID,
		DriverID:          trip.DriverID,
		PassengerID:       trip.PassengerID,
		TripStatus:        trip.TripStatus,
		TripType:          trip.TripType,
		PaymentMethod:     trip.PaymentMethod,
		PaymentStatus:     trip.PaymentStatus,
		PickupTime:        trip.PickupTime,
		DropTime:          trip.DropTime,
		PickupAddress:     trip.PickupAddress,
		PickupLat:         trip.PickupLat,
		PickupLon:         trip.PickupLon,
		DropAddress:       trip.DropAddress,
		DropLat:           trip.DropLat,
		DropLon:           trip.DropLon,
		DistanceTravelled: trip.DistanceTravelled,
		TotalFare:         trip.TotalFare,
		Discount:          trip.Discount,
		FinalFare:         trip.FinalFare,
	}

	tripID, err := t.TripRepository.Save(tripModel)
	if err != nil {
		return 0, err // Return error if save operation failed
	}

	return tripID, nil
}

func (t *TripServiceImpl) UpdateTravelStatus(trips request.UpdateTripStatusRequest) error {
	err := t.Validate.Struct(trips)
	if err != nil {
		return err
	}

	tripData := model.Trips{
		TripID:     trips.TripID,
		TripStatus: trips.TripStatus,
	}
	err = t.TripRepository.UpdateTravelStatus(tripData)
	if err != nil {
		return err
	}

	return nil
}
