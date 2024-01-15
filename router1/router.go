package router1

import (
	"TripManagementSystem/config"
	"TripManagementSystem/controller"
	"TripManagementSystem/helper"
	"TripManagementSystem/repository"
	"TripManagementSystem/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Dependencies struct {
	// TagsController      *controller.TagsController
	VehicleController   *controller.VehicleController
	DriverController    *controller.DriverController
	TripController      *controller.TripController
	DVMappingController *controller.DVMappingController
}

func InitializeDependencies() *Dependencies {
	db := config.DatabaseConnection()
	validate := validator.New()

	// Call AutoMigrate from migrations.go
	if err := helper.AutoMigrate(db); err != nil {
		// Handle error, for example, log it or panic
		log.Fatalf("Migration failed: %v", err)
	}

	// Setup for Vehicle
	vehicleRepository := repository.NewVehicleRepositoryImpl(db)
	vehicleService := service.NewVehicleServiceImpl(vehicleRepository, validate)
	vehicleController := controller.NewVehicleController(vehicleService)

	// Setup for Driver
	driverRepository := repository.NewDriverRepositoryImpl(db)
	driverService := service.NewDriverServiceImpl(driverRepository, validate)
	driverController := controller.NewDriverController(driverService)

	// Setup for Driver
	tripRepository := repository.NewTripRepositoryImpl(db)
	tripService := service.NewTripServiceImpl(tripRepository, validate)
	tripController := controller.NewTripController(tripService)

	// Add Vehicle Driver Mapping
	dvMappingRepository := repository.NewDVMappingRepositoryImpl(db)
	dvMappingService := service.NewDVMappingServiceImpl(dvMappingRepository, validate)
	dvMappingController := controller.NewDVMappingController(dvMappingService)

	return &Dependencies{
		// TagsController:      tagsController,
		VehicleController:   vehicleController,
		DriverController:    driverController,
		TripController:      tripController,
		DVMappingController: dvMappingController,
	}
}

func NewRouter(deps *Dependencies) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome home")
	})

	baseRouter := router.Group("/v1")

	vehicleRouter := baseRouter.Group("/vehicle")
	vehicleRouter.POST("", deps.VehicleController.Create)

	driverRouter := baseRouter.Group("/driver")
	driverRouter.POST("", deps.DriverController.Create)

	tripRouter := baseRouter.Group("/trip")
	tripRouter.POST("", deps.TripController.Create)
	tripRouter.PUT("/status/:tripId", deps.TripController.UpdateTripStatus)

	dvMappingRouter := baseRouter.Group("/dv_mapping")
	dvMappingRouter.POST("", deps.DVMappingController.Create)

	return router
}
