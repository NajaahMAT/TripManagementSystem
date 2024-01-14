package model

import (
	"time"
)

// Trip represents the details of a trip in the system.
type Trips struct {
	TripID            int64     `gorm:"primary_key;auto_increment;comment:'Unique identifier for the trip'"`
	VehicleID         int64     `gorm:"not null;comment:'Reference to the vehicle used in the trip'"`
	DriverID          int64     `gorm:"not null;comment:'Reference to the driver of the trip'"`
	PassengerID       int64     `gorm:"not null;comment:'Reference to the passenger of the trip'"`
	TripStatus        string    `gorm:"type:varchar(50);comment:'Current status of the trip'"`
	TripType          string    `gorm:"type:varchar(50);comment:'Type of the trip, Eg: Ride/Package/Rental'"`
	PaymentMethod     int       `gorm:"type:int;comment:'Identifier for the payment method used'"`
	PaymentStatus     int       `gorm:"type:int;comment:'Status of the payment for the trip'"`
	PickupTime        time.Time `gorm:"comment:'Time when the passenger was picked up'"`
	DropTime          time.Time `gorm:"comment:'Time when the passenger was dropped off'"`
	PickupAddress     string    `gorm:"type:varchar(255);comment:'Pickup address of the passenger'"`
	PickupLat         float64   `gorm:"type:decimal(10,6);comment:'Latitude of the pickup location'"`
	PickupLon         float64   `gorm:"type:decimal(10,6);comment:'Longitude of the pickup location'"`
	DropAddress       string    `gorm:"type:varchar(255);comment:'Drop address of the passenger'"`
	DropLat           float64   `gorm:"type:decimal(10,6);comment:'Latitude of the drop location'"`
	DropLon           float64   `gorm:"type:decimal(10,6);comment:'Longitude of the drop location'"`
	DistanceTravelled float64   `gorm:"type:decimal(10,2);comment:'Total distance travelled during the trip in kilometers'"`
	TotalFare         float64   `gorm:"type:decimal(10,2);comment:'Total fare calculated for the trip'"`
	Discount          float64   `gorm:"type:decimal(10,2);comment:'Discount applied on the trip fare'"`
	FinalFare         float64   `gorm:"type:decimal(10,2);comment:'Final fare amount after discounts'"`
	CreatedDate       time.Time `gorm:"autoCreateTime;comment:'Timestamp when the trip record was created'"`
	UpdatedDate       time.Time `gorm:"autoUpdateTime;comment:'Timestamp when the trip record was last updated'"`
	Vehicle           Vehicles  `gorm:"foreignKey:VehicleID;references:VehicleID"`
	Driver            Drivers   `gorm:"foreignKey:DriverID;references:DriverID"`
}
