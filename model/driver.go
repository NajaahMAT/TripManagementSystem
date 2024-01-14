package model

import "time"

type Drivers struct {
	DriverID      int64     `gorm:"primary_key;auto_increment;comment:'Unique identifier for the driver'"`
	FirstName     string    `gorm:"type:varchar(50)"`
	LastName      string    `gorm:"type:varchar(50)"`
	Surname       string    `gorm:"type:varchar(50)"`
	LicenseNumber string    `gorm:"type:varchar(50);not null;unique;comment:'Drivers license number'"`
	Email         string    `gorm:"type:varchar(50);comment:'Email address of the driver'"`
	Gender        string    `gorm:"type:varchar(50);comment:'Male/Female'"`
	Dob           string    `gorm:"type:varchar(50);comment:'Date of Birth of the driver'"`
	MobileNo      string    `gorm:"type:varchar(50);comment:'Drivers working mobile number'"`
	Address       string    `gorm:"type:varchar(50);comment:'Current Address of the driver'"`
	Language      string    `gorm:"type:varchar(50);comment:'Primary language of the driver'"`
	AccountNumber string    `gorm:"type:varchar(50);comment:'Drivers bank account number'"`
	DeviceID      string    `gorm:"type:varchar(50);comment:'Driver Device ID'"`
	DeviceType    string    `gorm:"type:varchar(50);comment:'Driver Device Type'"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}
