package request

import "time"

type CreateTripRequest struct {
	VehicleID         int64     `json:"vehicle_id" validate:"required"`
	DriverID          int64     `json:"driver_id" validate:"required"`
	PassengerID       int64     `json:"passenger_id" validate:"required"`
	TripStatus        string    `json:"trip_status"`
	TripType          string    `json:"trip_type" validate:"required"`
	PaymentMethod     int       `json:"payment_method" validate:"required"`
	PaymentStatus     int       `json:"payment_status"`
	PickupTime        time.Time `json:"pickup_time"`
	DropTime          time.Time `json:"drop_time"`
	PickupAddress     string    `json:"pickup_address" validate:"required"`
	PickupLat         float64   `json:"pickup_lat" validate:"required"`
	PickupLon         float64   `json:"pickup_lon" validate:"required"`
	DropAddress       string    `json:"drop_address"`
	DropLat           float64   `json:"drop_lat"`
	DropLon           float64   `json:"drop_lon"`
	DistanceTravelled float64   `json:"distance_travelled"`
	TotalFare         float64   `json:"total_fare"`
	Discount          float64   `json:"discount"`
	FinalFare         float64   `json:"final_fare"`
}

// TripStatus defines the possible states of a Trip.
type TripStatus string

const (
	TripStatusCreated TripStatus = "Created"
	// TripStatusDriverAssigned TripStatus = "DriverAssigned"
	// TripStatusDriverAccepted TripStatus = "DriverAccepted"
	// TripStatusDriverRejected TripStatus = "DriverRejected"
	// TripStatusDriverArrived  TripStatus = "DriverArrived"
	// TripStatusStarted    TripStatus = "Started"
	TripStatusInProgress TripStatus = "InProgress"
	// TripStatusEnded          TripStatus = "Ended"
	TripStatusCompleted TripStatus = "Completed"
	// TripStatusCancelled      TripStatus = "Cancelled"
)
