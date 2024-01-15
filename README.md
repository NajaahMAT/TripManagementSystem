To Run the Project:
1. Clone the go project. using the following command:   `git clone https://github.com/NajaahMAT/TripManagementSystem.git`
2. Check out `master` branch.  `git checkout master
3. Update go modules.  `go mod tidy`
4. Establish the Database Connection , Provide the following Params 
    ```
    const (
        user     = "root"       // change as per your MySQL user
        password = "password" // change as per your MySQL password
        dbName   = "trip_management"
        host     = "localhost"
        port     = 3306
    )
    ```
Database Configurations Should be edited in the "TripManagementSystem/config/mysql_database.go" file
5. To Build the project.  `go build`
6. To execute the project. `go run .`

# Trip Management System
------------------------

The system will need to store information about vehicles, drivers and trips records. Additionally, the application should be designed to handle concurrent requests efficiently, ensuring data consistency and optimal performance. Assumptions: a.vehicle can have many drivers b.trip can be created with one driver and one vehicle.

Assumptions:
a.vehicle can have many drivers
b.trip can be created with one driver and one vehicle

=========================================================
Database Design:
a. Outline the database schema, including tables, relationships, and key constraints.
Database Configurations Should be edited in the "TripManagementSystem/config/mysql_database.go" file
```
const (
	user     = "root"       // change as per your MySQL user
	password = "password" // change as per your MySQL password
	dbName   = "trip_management"
	host     = "localhost"
	port     = 3306
)
```

The Tables , Relationships and Key constraints are created through AutoMigration.For the file path "TripManagementSystem/helper/db_migrations.go"
Here we have created 4 Tables.
1. To Store Driver Data:  `drivers`
2. To Store Vehicle Data: `vehicles`
3. To Store Trip Data: `trips`
4. As per the assumptions mentioned in the question "vehicle can have many drivers"  , table "driver_vehicle_mappings" created. Once the drivers and vehicles created drivers will be mapped to the vehicle.
Following is the Create Table Queries associated with each tables:

***  Create `drivers` table
```
CREATE TABLE `drivers` (
  `driver_id` bigint NOT NULL AUTO_INCREMENT COMMENT '''Unique identifier for the driver''',
  `first_name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `surname` varchar(50) DEFAULT NULL,
  `license_number` varchar(50) NOT NULL COMMENT '''Drivers license number''',
  `email` varchar(50) DEFAULT NULL COMMENT '''Email address of the driver''',
  `gender` varchar(50) DEFAULT NULL COMMENT '''Male/Female''',
  `dob` varchar(50) DEFAULT NULL COMMENT '''Date of Birth of the driver''',
  `mobile_no` varchar(50) DEFAULT NULL COMMENT '''Drivers working mobile number''',
  `address` varchar(50) DEFAULT NULL COMMENT '''Current Address of the driver''',
  `language` varchar(50) DEFAULT NULL COMMENT '''Primary language of the driver''',
  `account_number` varchar(50) DEFAULT NULL COMMENT '''Drivers bank account number''',
  `device_id` varchar(50) DEFAULT NULL COMMENT '''Driver Device ID''',
  `device_type` varchar(50) DEFAULT NULL COMMENT '''Driver Device Type''',
  `created_date` datetime(3) DEFAULT NULL,
  `updated_date` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`driver_id`),
  UNIQUE KEY `license_number` (`license_number`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```
| Column Name     | Data Type       | Constraints                               | Description                               |
|-----------------|-----------------|-------------------------------------------|-------------------------------------------|
| `driver_id`     | bigint          | PRIMARY KEY, NOT NULL, AUTO_INCREMENT     | Unique identifier for the driver          |
| `first_name`    | varchar(50)     | DEFAULT NULL                              |                                           |
| `last_name`     | varchar(50)     | DEFAULT NULL                              |                                           |
| `surname`       | varchar(50)     | DEFAULT NULL                              |                                           |
| `license_number`| varchar(50)     | NOT NULL, UNIQUE KEY                      | Driver's license number                   |
| `email`         | varchar(50)     | DEFAULT NULL                              | Email address of the driver               |
| `gender`        | varchar(50)     | DEFAULT NULL                              | Male/Female                               |
| `dob`           | varchar(50)     | DEFAULT NULL                              | Date of Birth of the driver               |
| `mobile_no`     | varchar(50)     | DEFAULT NULL                              | Driver's working mobile number            |
| `address`       | varchar(50)     | DEFAULT NULL                              | Current Address of the driver             |
| `language`      | varchar(50)     | DEFAULT NULL                              | Primary language of the driver            |
| `account_number`| varchar(50)     | DEFAULT NULL                              | Driver's bank account number              |
| `device_id`     | varchar(50)     | DEFAULT NULL                              | Driver Device ID                          |
| `device_type`   | varchar(50)     | DEFAULT NULL                              | Driver Device Type                        |
| `created_date`  | datetime(3)     | DEFAULT NULL                              |                                           |
| `updated_date`  | datetime(3)     | DEFAULT NULL                              |                                           |


***  Create `vehicles` table
```
CREATE TABLE `vehicles` (
  `vehicle_id` bigint NOT NULL AUTO_INCREMENT COMMENT '''Unique identifier for the vehicle''',
  `type` varchar(50) DEFAULT NULL COMMENT '''The type of the vehicle, Example:Car/Van/Bus''',
  `registration_number` varchar(50) NOT NULL COMMENT '''Vehicle''s registration number''',
  `model_make` varchar(50) DEFAULT NULL COMMENT '''Make of the vehicle''s model''',
  `brand_model` varchar(50) DEFAULT NULL COMMENT '''Brand Model of the vehicle''',
  `color` varchar(50) DEFAULT NULL COMMENT '''Color of the vehicle''',
  `year` year DEFAULT NULL COMMENT '''Manufacturing year of the vehicle''',
  `seating_capacity` bigint DEFAULT NULL COMMENT '''Seating capacity of the vehicle''',
  `engine_capacity` bigint DEFAULT NULL COMMENT '''Engine capacity of the vehicle''',
  `created_date` datetime(3) DEFAULT NULL,
  `updated_date` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`vehicle_id`),
  UNIQUE KEY `registration_number` (`registration_number`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```

| Column Name          | Data Type        | Key Constraint    | Description                                  |
|----------------------|------------------|-------------------|----------------------------------------------|
| `vehicle_id`         | bigint           | PRIMARY KEY       | Unique identifier for the vehicle            |
| `type`               | varchar(50)      |                   | The type of the vehicle, e.g., Car/Van/Bus   |
| `registration_number`| varchar(50)      | UNIQUE KEY        | Vehicle's registration number                |
| `model_make`         | varchar(50)      |                   | Make of the vehicle's model                  |
| `brand_model`        | varchar(50)      |                   | Brand Model of the vehicle                   |
| `color`              | varchar(50)      |                   | Color of the vehicle                         |
| `year`               | year             |                   | Manufacturing year of the vehicle            |
| `seating_capacity`   | bigint           |                   | Seating capacity of the vehicle              |
| `engine_capacity`    | bigint           |                   | Engine capacity of the vehicle               |
| `created_date`       | datetime(3)      |                   | Timestamp when the record was created        |
| `updated_date`       | datetime(3)      |                   | Timestamp when the record was last updated   |


***  Create `trips` table
```
CREATE TABLE `trips` (
  `trip_id` bigint NOT NULL AUTO_INCREMENT COMMENT '''Unique identifier for the trip''',
  `vehicle_id` bigint NOT NULL COMMENT '''Reference to the vehicle used in the trip''',
  `driver_id` bigint NOT NULL COMMENT '''Reference to the driver of the trip''',
  `passenger_id` bigint NOT NULL COMMENT '''Reference to the passenger of the trip''',
  `trip_status` varchar(50) DEFAULT NULL COMMENT '''Current status of the trip''',
  `trip_type` varchar(50) DEFAULT NULL COMMENT '''Type of the trip, Eg: Ride/Package/Rental''',
  `payment_method` bigint DEFAULT NULL COMMENT '''Identifier for the payment method used''',
  `payment_status` bigint DEFAULT NULL COMMENT '''Status of the payment for the trip''',
  `pickup_time` datetime(3) DEFAULT NULL COMMENT '''Time when the passenger was picked up''',
  `drop_time` datetime(3) DEFAULT NULL COMMENT '''Time when the passenger was dropped off''',
  `pickup_address` varchar(255) DEFAULT NULL COMMENT '''Pickup address of the passenger''',
  `pickup_lat` decimal(10,6) DEFAULT NULL COMMENT '''Latitude of the pickup location''',
  `pickup_lon` decimal(10,6) DEFAULT NULL COMMENT '''Longitude of the pickup location''',
  `drop_address` varchar(255) DEFAULT NULL COMMENT '''Drop address of the passenger''',
  `drop_lat` decimal(10,6) DEFAULT NULL COMMENT '''Latitude of the drop location''',
  `drop_lon` decimal(10,6) DEFAULT NULL COMMENT '''Longitude of the drop location''',
  `distance_travelled` decimal(10,2) DEFAULT NULL COMMENT '''Total distance travelled during the trip in kilometers''',
  `total_fare` decimal(10,2) DEFAULT NULL COMMENT '''Total fare calculated for the trip''',
  `discount` decimal(10,2) DEFAULT NULL COMMENT '''Discount applied on the trip fare''',
  `final_fare` decimal(10,2) DEFAULT NULL COMMENT '''Final fare amount after discounts''',
  `created_date` datetime(3) DEFAULT NULL COMMENT '''Timestamp when the trip record was created''',
  `updated_date` datetime(3) DEFAULT NULL COMMENT '''Timestamp when the trip record was last updated''',
  PRIMARY KEY (`trip_id`),
  KEY `fk_vehicles` (`vehicle_id`),
  KEY `fk_drivers` (`driver_id`),
  CONSTRAINT `fk_drivers` FOREIGN KEY (`driver_id`) REFERENCES `drivers` (`driver_id`),
  CONSTRAINT `fk_vehicles` FOREIGN KEY (`vehicle_id`) REFERENCES `vehicles` (`vehicle_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```
Note: In here instead of using `driver_id` and `trip_id` seperately we can use the `driver_vehicle_mappings.mapping_id` to create the trip. 

| Column Name         | Data Type        | Constraints                      | Description                                                |
|---------------------|------------------|----------------------------------|------------------------------------------------------------|
| `trip_id`           | bigint           | NOT NULL, PRIMARY KEY, AUTO_INCREMENT | Unique identifier for the trip                            |
| `vehicle_id`        | bigint           | NOT NULL, FOREIGN KEY            | Reference to the vehicle used in the trip                 |
| `driver_id`         | bigint           | NOT NULL, FOREIGN KEY            | Reference to the driver of the trip                       |
| `passenger_id`      | bigint           | NOT NULL                         | Reference to the passenger of the trip                    |
| `trip_status`       | varchar(50)      | DEFAULT NULL                     | Current status of the trip                                |
| `trip_type`         | varchar(50)      | DEFAULT NULL                     | Type of the trip (e.g., Ride/Package/Rental)              |
| `payment_method`    | bigint           | DEFAULT NULL                     | Identifier for the payment method used                    |
| `payment_status`    | bigint           | DEFAULT NULL                     | Status of the payment for the trip                        |
| `pickup_time`       | datetime(3)      | DEFAULT NULL                     | Time when the passenger was picked up                     |
| `drop_time`         | datetime(3)      | DEFAULT NULL                     | Time when the passenger was dropped off                   |
| `pickup_address`    | varchar(255)     | DEFAULT NULL                     | Pickup address of the passenger                           |
| `pickup_lat`        | decimal(10,6)    | DEFAULT NULL                     | Latitude of the pickup location                           |
| `pickup_lon`        | decimal(10,6)    | DEFAULT NULL                     | Longitude of the pickup location                          |
| `drop_address`      | varchar(255)     | DEFAULT NULL                     | Drop address of the passenger                             |
| `drop_lat`          | decimal(10,6)    | DEFAULT NULL                     | Latitude of the drop location                             |
| `drop_lon`          | decimal(10,6)    | DEFAULT NULL                     | Longitude of the drop location                            |
| `distance_travelled`| decimal(10,2)    | DEFAULT NULL                     | Total distance travelled during the trip in kilometers   |
| `total_fare`        | decimal(10,2)    | DEFAULT NULL                     | Total fare calculated for the trip                        |
| `discount`          | decimal(10,2)    | DEFAULT NULL                     | Discount applied on the trip fare                         |
| `final_fare`        | decimal(10,2)    | DEFAULT NULL                     | Final fare amount after discounts                         |
| `created_date`      | datetime(3)      | DEFAULT NULL                     | Timestamp when the trip record was created                |
| `updated_date`      | datetime(3)      | DEFAULT NULL                     | Timestamp when the trip record was last updated           |

- `fk_vehicles`: FOREIGN KEY (`vehicle_id`) REFERENCES `vehicles


***  Create `drver_vehicle_mappings` table
```
CREATE TABLE `driver_vehicle_mappings` (
  `mapping_id` bigint NOT NULL AUTO_INCREMENT COMMENT '''Unique identifier for the DriverVehicleMapping''',
  `driver_id` bigint NOT NULL COMMENT '''Reference to the driver''',
  `vehicle_id` bigint NOT NULL COMMENT '''Reference to the vehicle''',
  `started_at` datetime(3) DEFAULT NULL COMMENT '''Time when the mapping affected from''',
  `ended_at` datetime(3) DEFAULT NULL COMMENT '''Time when the mapping ended on''',
  `created_date` datetime(3) DEFAULT NULL,
  `updated_date` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`mapping_id`),
  KEY `fk_vehicles1` (`vehicle_id`),
  KEY `fk_drivers2` (`driver_id`),
  CONSTRAINT `fk_drivers2` FOREIGN KEY (`driver_id`) REFERENCES `drivers` (`driver_id`),
  CONSTRAINT `fk_vehicles1` FOREIGN KEY (`vehicle_id`) REFERENCES `vehicles` (`vehicle_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```
| Column Name   | Data Type        | Constraints                          | Comments                                  |
|---------------|------------------|--------------------------------------|-------------------------------------------|
| mapping_id    | bigint           | PRIMARY KEY, NOT NULL, AUTO_INCREMENT| 'Unique identifier for the DriverVehicleMapping' |
| driver_id     | bigint           | NOT NULL, FOREIGN KEY                | 'Reference to the driver'                 |
| vehicle_id    | bigint           | NOT NULL, FOREIGN KEY                | 'Reference to the vehicle'                |
| started_at    | datetime(3)      | DEFAULT NULL                         | 'Time when the mapping affected from'     |
| ended_at      | datetime(3)      | DEFAULT NULL                         | 'Time when the mapping ended on'          |
| created_date  | datetime(3)      | DEFAULT NULL                         |                                           |
| updated_date  | datetime(3)      | DEFAULT NULL                         |                                           |

- The `PRIMARY KEY` is defined on the `mapping_id` column.
- The `driver_id` column references the `driver_id` in the `drivers` table (`fk_drivers2`).
- The `vehicle_id` column references the `vehicle_id` in the `vehicles` table (`fk_vehicles1`).

=========================================================
b. Justify your design choices, considering factors such as scalability, data integrity, and ease of
querying.

*** Relationships
A vehicle can be driven by many drivers over time, and a driver can drive many vehicles. This many-to-many relationship is captured in the `driver_vehicle_mappings` table.
Each trip is associated with exactly one vehicle and one driver, creating a one-to-one relationship between the Trips table and both the Vehicles and Drivers tables.

*** Constraints and Indexes
Use of Primary and Foreign Key Constraints Ensure the data integrity by linking tables appropriately.
Unique Constraints: On fields like `registration_number` in `vehicles` table, and `license_number` in `drivers` table, to avoid duplicates.
Indexes: On frequently queried columns like `driver_id` , `license_number` in `driver` table  and `vehicle_id`, `license_number` in the `vehicles` table and `trip_id`, `driver_id`,`vehicle_id` in the `trips` table and `mapping_id`, `driver_id`, `vehicle_id` in `driver_vehicle_mappings` table indexed for faster searches.

***Considerations for Scalability and Performance
Scalability and Performance of the Database Schema considering the Normalization of the database normalized to reduce redundancy and improve data integrity and Indexing Strategy used on query patterns to speed up data retrieval and the partitioning larger tables (like Trips) based on criteria like date or status for faster query performance in a large dataset.

## 1. Drivers Table
- **Normalization**: Up to 2NF as there are no partial dependencies on a composite key.
- **Indexing**: Primary key on `driver_id` and unique key on `license_number` for fast lookup.
- **Partitioning**: Could be partitioned based on `created_date` for better performance.

## 2. Vehicles Table
- **Normalization**: Up to 2NF as each non-key attribute is fully functionally dependent on the primary key.
- **Indexing**: Primary key on `vehicle_id` and unique key on `registration_number` enhance search efficiency.
- **Partitioning**: Potential for partitioning on attributes like `year` or `created_date`.

## 3. Trips Table
- **Normalization**: Up to 3NF as all attributes are dependent on the primary key and there is no transitive dependency.
- **Indexing**: Indexing on vehicle_id and driver_id for efficient joins and lookups. Additional indexing could be added for trip_status and trip_type for query optimization.
- **Partitioning**: Potential for partitioning based on pickup_time or created_date for handling large datasets efficiently.

## 4.Driver Vehicle Mappings Table
- **Normalization**: Up to 3NF as all attributes are dependent on the primary key without transitive dependencies.
- **Indexing**: Primary key indexing on mapping_id and foreign keys on driver_id and vehicle_id for fast data
retrieval and relationship integrity.
- **Partitioning**: Could be partitioned based on started_at or ended_at to manage historical data more efficiently.

=========================================================
Concurrency in Golang:
a. Implement a Golang program that simulates concurrent requests to the Trip Management System.
The program should perform the following operations concurrently:
    - Add a new vehicle record to the database.
    - Add a new driver record to the database.
    - Add a new trip record to the database.
    - Update the status of a trip to "in progress"(Assume it has created,in-progress and completed status).

#Overview of the Trip Management System API
BaseUrl: http://localhost:8888/v1/
PostmanCollection: path: "TripManagementSystem/postman_api_collection.json"
Api Endpoints: 
1. Add New Vehicle 
2. Add New Driver
3. Mapping Driver with Vehicle
4. Add New Trip
5. Update the Travel Status


# HTTP Requests and Responses

# HTTP Requests and Responses

| Name              | Endpoint               | HTTP Method | Content Type             | Request                                                                                                                                                                                                                                                                                                                                                                                                          | Response                                                     |
|-------------------|------------------------|-------------|--------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------|
| Add Vehicle       | `/v1/vehicle`          | POST        | `application/json`       | `{ "RegistrationNumber": "XYZ900005", "Type": "Car", "ModelMake": "Nissan", "BrandModel": "Leaf", "Color": "Blue", "Year": 2023, "SeatingCapacity": 6, "EngineCapacity": 3600 }`                                                                                                                                                                                                                                   | `{ "code": 200, "status": "Ok", "data": { "vehicle_id": 1 } }`   |
| Add Driver        | `/v1/driver`           | POST        | `application/json`       | `{ "firstName": "Joe", "lastName": "Wan", "surname": "Sha", "licenseNumber": "XYZ90000", "email": "joe@example.com", "gender": "Female", "dob": "1985-01-23", "mobileNo": "+1234567890", "address": "123, Baker Street, London", "language": "English", "accountNumber": "ACC1234567", "deviceId": "DEV12345", "deviceType": "Smartphone" }`                                                                       | `{ "code": 200, "status": "Ok", "data": { "driver_id": 1 } }`   |
| Add Trip          | `/v1/trip`             | POST        | `application/json`       | `{ "vehicle_id": 1, "driver_id": 1, "trip_type": "Ride", "passenger_id": 129, "trip_status": "Created", "payment_method": 1, "payment_status": 0, "pickup_time": "2022-01-01T15:04:05Z", "drop_time": "2022-01-01T16:04:05Z", "pickup_address": "123 Pickup St", "pickup_lat": 40.7128, "pickup_lon": -74.0060, "drop_address": "", "drop_lat": 0, "drop_lon": 0, "distance_travelled": 0, "total_fare": 0, "discount": 0, "final_fare": 0 }` | `{ "code": 200, "status": "Ok", "data": { "trip_id": 1 } }`     |
| Update Trip Status | `/v1/trip/status/1`    | PUT         | `application/json`       | `{ "trip_status": "DriverAssigned" }`                                                                                                                                                                                                                                                                                                                                                                             | `{ "code": 200, "status": "Ok" }`                                |
| Add DV Mapping    | `/v1/dv_mapping`       | POST        | `application/json`       | `{ "driver_id": 1, "vehicle_id": 1, "mapping_started_at": "2024-01-12T15:04:05Z", "mapping_ended_at": "2024-01-13T15:04:05Z" }`                                                                                                                                                                                                                                                                                   | `{ "code": 200, "status": "Ok", "data": { "mapping_id": 1 } }`  |

=========================================================
b. Address potential race conditions and ensure data consistency.
* Updating Trip Status of the trip:    When the trip status is being read and updated simutaneously by two requests it can cause the race condition. Transactions are crucial for ensuring data consistency, especially when dealing with operations that should be atomic. Here we have to wrap the read and write operations in a transaction. So one will have to wait for the other to complete, ensuring data integrity. 

Also we can use Transactions or use Mutexes for synchronization(sync.Mutex) to ensure only one goroutine accesses a shared resource at a time.This prevents race conditions since a mutex locks the resource, blocking other goroutines until the current goroutine finishes its operation and unlocks the mutex.

More than the above solutions we can do consistent error handling.In concurrent operations, consistently handle errors. If an error occurs in one goroutine, it should not leave the system in a half-updated state.In case of errors, use rollbacks or compensating transactions to revert changes and maintain data consistency


Following is the way race condition is handled for trip status update in this project.
```
func (t *TripRepositoryImpl) UpdateTravelStatus(trip model.Trips) error {
	tx := t.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// Fetch the current state of the trip to check its current status
	var currentTrip model.Trips
	if err := tx.Where("trip_id = ?", trip.TripID).First(&currentTrip).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Check if the trip is in a state that allows status update
	if (currentTrip.TripStatus != "Created" && trip.TripStatus == "InProgress") ||
		(currentTrip.TripStatus != "InProgress" && trip.TripStatus == "Completed") {
		tx.Rollback()
		return errors.New("trip status cannot be updated from its current state")
	}

	// Proceed with updating the status
	if err := tx.Model(&model.Trips{}).Where("trip_id = ?", trip.TripID).Update("trip_status", trip.TripStatus).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
```

=========================================================
c. Use appropriate Golang concurrency primitives (e.g., goroutines, channels) to achieve parallelism.
This project designed as APIs to perform the Operations, To acheive parallelism I have implemented it as a unit test.
path: "TripManagementSystem/concurrent_operations_test.go"
```
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
		FirstName:     "Jack",
		LastName:      "Jaan",
		Surname:       "la",
		LicenseNumber: "XYZ90006",
		Email:         "jacky.jaan@abc.com",
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
		RegistrationNumber: "Reg90005",
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
```

Test Result:
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestConcurrentOperations$ TripManagementSystem

=== RUN   TestConcurrentOperations

2024/01/14 21:49:38 c:/Users/LENOVO/go/src/TripManagementSystem/repository/vehicle_repository_impl.go:29 SLOW SQL >= 200ms
[250.009ms] [rows:1] INSERT INTO `vehicles` (`type`,`registration_number`,`model_make`,`brand_model`,`color`,`year`,`seating_capacity`,`engine_capacity`,`created_date`,`updated_date`) VALUES ('Car','Reg90005','Suzuki','Swift','Blue',2023,5,1800,'2024-01-14 21:49:38.474','2024-01-14 21:49:38.474')
    c:\Users\LENOVO\go\src\TripManagementSystem\concurrent_operations_test.go:79: Vehicle added with ID: 2

2024/01/14 21:49:38 c:/Users/LENOVO/go/src/TripManagementSystem/repository/driver_repository_impl.go:29 SLOW SQL >= 200ms
[278.043ms] [rows:1] INSERT INTO `drivers` (`first_name`,`last_name`,`surname`,`license_number`,`email`,`gender`,`dob`,`mobile_no`,`address`,`language`,`account_number`,`device_id`,`device_type`,`created_date`,`updated_date`) VALUES ('Jack','Jaan','la','XYZ90006','jacky.jaan@abc.com','Male','1995-01-08','+01187967850','1234, Armour Street, Dubai','English','ACC1232345','D4567876','Apple','2024-01-14 21:49:38.474','2024-01-14 21:49:38.474')
    c:\Users\LENOVO\go\src\TripManagementSystem\concurrent_operations_test.go:55: Driver added with ID: 2

2024/01/14 21:49:39 c:/Users/LENOVO/go/src/TripManagementSystem/repository/trip_repository_impl.go:29 SLOW SQL >= 200ms
[271.525ms] [rows:1] INSERT INTO `trips` (`vehicle_id`,`driver_id`,`passenger_id`,`trip_status`,`trip_type`,`payment_method`,`payment_status`,`pickup_time`,`drop_time`,`pickup_address`,`pickup_lat`,`pickup_lon`,`drop_address`,`drop_lat`,`drop_lon`,`distance_travelled`,`total_fare`,`discount`,`final_fare`,`created_date`,`updated_date`) VALUES (2,2,234,'','Ride',1,0,'2024-01-14 21:49:38.75','2024-01-14 21:49:38.75','123, abc street, xyz',4.899,78.9,'',0,0,0,0,0,0,'2024-01-14 21:49:38.755','2024-01-14 21:49:38.755')
    c:\Users\LENOVO\go\src\TripManagementSystem\concurrent_operations_test.go:109: Trip added with ID: 3
    c:\Users\LENOVO\go\src\TripManagementSystem\concurrent_operations_test.go:121: Trip 3 updated to Created
    c:\Users\LENOVO\go\src\TripManagementSystem\concurrent_operations_test.go:121: Trip 3 updated to InProgress
    c:\Users\LENOVO\go\src\TripManagementSystem\concurrent_operations_test.go:121: Trip 3 updated to Completed
--- PASS: TestConcurrentOperations (3.82s)
PASS
ok      TripManagementSystem    5.391s

```
=========================================================
Advanced Concepts:
a. Explain how you would handle long-running tasks (e.g., complex database queries, external API calls) concurrently without blocking the main application flow.
# Use of Goroutines for Asynchronous Execution
   * Launch long-running tasks such as complex database queries or external API calls in separate goroutines. This approach allows these tasks to run concurrently without blocking the main application flow.
   * Example: go performComplexQuery(db)

# Channels for Result Communication
   * Use channels to communicate the results of long-running tasks back to the main routine or other parts of the application. This way, the main application flow can continue executing and only process the results when ready.
   * Example: A channel can be used to send the result of a database query back to the main routine.

# Context Package for Timeout and Cancellation
   * Utilize the context package to set deadlines or timeouts for tasks. This is particularly useful for external API calls where response times can be unpredictable.
   * Example: ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

# WaitGroups for Synchronization
   * In scenarios where the main routine needs to wait for all concurrent tasks to complete, use sync.WaitGroup. This ensures that the application only proceeds after the completion of all long-running tasks.
   * Example: wg.Add(1); go func() { defer wg.Done(); performTask() }()

# Select Statement for Handling Multiple Channels
   * When dealing with multiple concurrent tasks, use the select statement to wait on multiple channels. This is a powerful feature for handling different outcomes (like success, error, timeout) in a non-blocking manner.
   * Example: Using select to wait on channels that return results or errors from several goroutines.



=====================================================================================================================================
b. Share any experience or examples where you have applied advanced concurrency concepts in Golang.
Experience:  Created Price files with the estimated fares using the (Pickup , Drop Locations (Lat, Lon)) for all vehicle types simultaneously which are available within the selected region in the PickMe Application.
* I have spawn a goroutine for each vehicle type to calculate the fare estimate concurrently. This is done using go func(vt VehicleType) { ... }(vehicleType).
* WaitGroup to Wait for Goroutines: The sync.WaitGroup is used to wait for all goroutines to finish their execution
* Channel for Collecting Results: A channel fareEstimates is used to collect fare estimates. from each goroutine. The channel's buffer size is set to the length of vehicleTypes to prevent blocking when goroutines send data to the channel.
*  Another goroutine is used to close the fareEstimates channel once all fare calculation goroutines have completed. This is important to avoid deadlocks and allow the main goroutine to exit the range loop over the channel.
* I have used a range loop over the fareEstimates channel to collect all fare estimates. This loop runs until the channel is closed.
* Finally, the collected fare estimates are passed to createPriceFile, which simulates the creation of a price file