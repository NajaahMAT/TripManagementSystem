package helper

import (
	"TripManagementSystem/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	// Perform auto-migration for all models
	err := db.Table("vehicles").AutoMigrate(&model.Vehicles{})
	if err != nil {
		return err
	}

	err = db.Table("drivers").AutoMigrate(&model.Drivers{})
	if err != nil {
		return err
	}

	err = db.Table("trips").AutoMigrate(&model.Trips{})
	if err != nil {
		return err
	}

	// Set up foreign key constraints for Trips table
	db.Exec("ALTER TABLE trips ADD CONSTRAINT fk_vehicles FOREIGN KEY (vehicle_id) REFERENCES vehicles(vehicle_id)")
	db.Exec("ALTER TABLE trips ADD CONSTRAINT fk_drivers FOREIGN KEY (driver_id) REFERENCES drivers(driver_id)")

	err = db.Table("driver_vehicle_mappings").AutoMigrate(&model.DriverVehicleMappings{})
	if err != nil {
		return err
	}

	// Set up foreign key constraints for DriverVehicleMappings table
	db.Exec("ALTER TABLE driver_vehicle_mappings ADD CONSTRAINT fk_vehicles1 FOREIGN KEY (vehicle_id) REFERENCES vehicles(vehicle_id)")
	db.Exec("ALTER TABLE driver_vehicle_mappings ADD CONSTRAINT fk_drivers2 FOREIGN KEY (driver_id) REFERENCES drivers(driver_id)")

	return nil
}
