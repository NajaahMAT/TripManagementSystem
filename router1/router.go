package router1

import (
	"TripManagementSystem/config"
	"TripManagementSystem/controller"
	"TripManagementSystem/model"
	"TripManagementSystem/repository"
	"TripManagementSystem/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Dependencies struct {
	TagsController      *controller.TagsController
	VehicleController   *controller.VehicleController
	DriverController    *controller.DriverController
	TripController      *controller.TripController
	DVMappingController *controller.DVMappingController
}

func InitializeDependencies() *Dependencies {
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})
	db.Table("vehicles").AutoMigrate(&model.Vehicles{})
	db.Table("drivers").AutoMigrate(&model.Drivers{})
	db.Table("trips").AutoMigrate(&model.Trips{})
	db.Exec("ALTER TABLE trips ADD CONSTRAINT fk_vehicles FOREIGN KEY (vehicle_id) REFERENCES vehicles(vehicle_id)")
	db.Exec("ALTER TABLE trips ADD CONSTRAINT fk_drivers FOREIGN KEY (driver_id) REFERENCES drivers(driver_id)")
	db.Table("driver_vehicle_mappings").AutoMigrate(&model.DriverVehicleMappings{})
	db.Exec("ALTER TABLE driver_vehicle_mappings ADD CONSTRAINT fk_vehicles1 FOREIGN KEY (vehicle_id) REFERENCES vehicles(vehicle_id)")
	db.Exec("ALTER TABLE driver_vehicle_mappings ADD CONSTRAINT fk_drivers2 FOREIGN KEY (driver_id) REFERENCES drivers(driver_id)")

	// Setup for tag
	tagsRepository := repository.NewTagsRepositoryImpl(db)
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
	tagsController := controller.NewTagsController(tagsService)

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
		TagsController:      tagsController,
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

	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", deps.TagsController.FindAll)
	tagsRouter.GET("/:tagsId", deps.TagsController.FindById)
	tagsRouter.POST("", deps.TagsController.Create)
	tagsRouter.PATCH("/:tagsId", deps.TagsController.Update)
	tagsRouter.DELETE("/:tagsId", deps.TagsController.Delete)

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
