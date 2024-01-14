package model

import "time"

type Vehicles struct {
	VehicleID          int64     `gorm:"primary_key;auto_increment;comment:'Unique identifier for the vehicle'"`
	Type               string    `gorm:"type:varchar(50);comment:'The type of the vehicle, Example:Car/Van/Bus'"`
	RegistrationNumber string    `gorm:"type:varchar(50);not null;unique;comment:'Vehicle's registration number'"`
	ModelMake          string    `gorm:"type:varchar(50);comment:'Make of the vehicle's model'"`
	BrandModel         string    `gorm:"type:varchar(50);comment:'Brand Model of the vehicle'"`
	Color              string    `gorm:"type:varchar(50);comment:'Color of the vehicle'"`
	Year               int       `gorm:"type:year;comment:'Manufacturing year of the vehicle'"`
	SeatingCapacity    int       `gorm:"type:int;comment:'Seating capacity of the vehicle'"`
	EngineCapacity     int       `gorm:"type:int;comment:'Engine capacity of the vehicle'"`
	CreatedDate        time.Time `gorm:"autoCreateTime"`
	UpdatedDate        time.Time `gorm:"autoUpdateTime"`
}
