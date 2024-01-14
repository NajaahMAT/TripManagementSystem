package request

type UpdateTripStatusRequest struct {
	TripID     int64  `validate:"required" json:"trip_id"`
	TripStatus string `validate:"required" json:"trip_status"`
}
