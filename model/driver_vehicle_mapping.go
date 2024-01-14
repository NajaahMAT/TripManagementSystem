package model

import "time"

type DriverVehicleMappings struct {
	MappingID   int64     `gorm:"primary_key;auto_increment;comment:'Unique identifier for the DriverVehicleMapping'"`
	DriverID    int64     `gorm:"not null;comment:'Reference to the driver'"`
	VehicleID   int64     `gorm:"not null;comment:'Reference to the vehicle'"`
	StartedAt   time.Time `gorm:"comment:'Time when the mapping affected from'"`
	EndedAt     time.Time `gorm:"comment:'Time when the mapping ended on'"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Vehicle     Vehicles  `gorm:"foreignKey:VehicleID;references:VehicleID"`
	Driver      Drivers   `gorm:"foreignKey:DriverID;references:DriverID"`
}
