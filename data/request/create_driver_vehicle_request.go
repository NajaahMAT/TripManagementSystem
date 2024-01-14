package request

import "time"

type CreateDVMappingRequest struct {
	DriverID         int64     `json:"driver_id"`
	VehicleID        int64     `json:"vehicle_id"`
	MappingStartedAt time.Time `json:"mapping_started_at"`
	MappingEndedAt   time.Time `json:"mapping_ended_at"`
}
