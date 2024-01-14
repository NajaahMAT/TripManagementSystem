package main

import (
	"TripManagementSystem/config"
	"TripManagementSystem/data/request"
	"TripManagementSystem/repository"
	"TripManagementSystem/service"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
)

// Mock or set up necessary dependencies
var (
	vehicleService = service.NewVehicleServiceImpl(repository.NewVehicleRepositoryImpl(config.DatabaseConnection()), validator.New())
	driverService  = service.NewDriverServiceImpl(repository.NewDriverRepositoryImpl(config.DatabaseConnection()), validator.New())
	tripService    = service.NewTripServiceImpl(repository.NewTripRepositoryImpl(config.DatabaseConnection()), validator.New())
)

func TestConcurrentOperations(t *testing.T) {
	var wg sync.WaitGroup

	errors := make(chan error, 4)
	driverIDChan := make(chan int64, 1)
	vehicleIDChan := make(chan int64, 1)

	mockDriverRequest := request.CreateDriverRequest{
		FirstName:     "Moon",
		LastName:      "Sun",
		Surname:       "la",
		LicenseNumber: "XYZ90001",
		Email:         "john.smith@gmail.com",
		Gender:        "Male",
		Dob:           "1995-01-08",
		MobileNo:      "+01187967850",
		Address:       "1234, Armour Street, Dubai",
		Language:      "English",
		AccountNumber: "ACC1232345",
		DeviceID:      "D4567876",
		DeviceType:    "Apple",
	}

	// Add a driver
	wg.Add(1)
	go func() {
		defer wg.Done()
		driverID, err := driverService.Create(mockDriverRequest)
		if err != nil {
			errors <- fmt.Errorf("add driver: %v", err)
		} else {
			driverIDChan <- driverID // Send driverID to the channel
			t.Logf("Driver added with ID: %d", driverID)
		}
	}()

	mockVehicleRequest := request.CreateVehicleRequest{
		RegistrationNumber: "Reg90001",
		Type:               "Car",
		ModelMake:          "Suzuki",
		BrandModel:         "Swift",
		Color:              "Blue",
		Year:               2023,
		SeatingCapacity:    5,
		EngineCapacity:     1800,
	}

	// Add a vehicle
	wg.Add(1)
	go func() {
		defer wg.Done()
		vehicleID, err := vehicleService.Create(mockVehicleRequest)
		if err != nil {
			errors <- fmt.Errorf("add vehicle: %v", err)
		} else {
			vehicleIDChan <- vehicleID // Send vehicleID to the channel
			t.Logf("Vehicle added with ID: %d", vehicleID)
		}
	}()

	// Add a trip and update its status
	wg.Add(1)
	go func() {
		defer wg.Done()
		driverID := <-driverIDChan   // Receive driverID from the channel
		vehicleID := <-vehicleIDChan // Receive vehicleID from the channel

		mockTripRequest := request.CreateTripRequest{
			VehicleID:     driverID,
			DriverID:      vehicleID,
			PassengerID:   234,
			TripStatus:    "",
			TripType:      "Ride",
			PaymentMethod: 1,
			PickupAddress: "123, abc street, xyz",
			PickupLat:     4.8990,
			PickupLon:     78.9,
			PickupTime:    time.Now(),
			DropTime:      time.Now(),
		}

		tripID, err := tripService.Create(mockTripRequest)
		if err != nil {
			errors <- fmt.Errorf("add trip: %v", err)
			return
		}
		t.Logf("Trip added with ID: %d", tripID)

		// Update trip status sequentially
		for _, status := range []string{"Created", "InProgress", "Completed"} {
			mockTripStatusRequest := request.UpdateTripStatusRequest{
				TripID:     tripID,
				TripStatus: status,
			}
			if err := tripService.UpdateTravelStatus(mockTripStatusRequest); err != nil {
				errors <- fmt.Errorf("update trip %d to %s: %v", tripID, status, err)
				return
			}
			t.Logf("Trip %d updated to %s", tripID, status)
			time.Sleep(1 * time.Second) // Simulate processing time
		}
	}()

	go func() {
		wg.Wait()
		close(errors)
		close(driverIDChan)
		close(vehicleIDChan)
	}()

	// Implement a timeout for the test
	timeout := time.After(10 * time.Second)

	for {
		select {
		case err := <-errors:
			if err != nil {

				t.Errorf("Operation failed: %v", err)
			}
			if len(errors) == 0 {
				return // Exit when all errors have been processed
			}
		case <-timeout:
			t.Fatal("Test timed out")
			return
		}
	}
}
