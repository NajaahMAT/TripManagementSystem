package controller

import (
	"TripManagementSystem/data/request"
	"TripManagementSystem/data/response"
	"TripManagementSystem/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	vehicleService service.VehicleService
}

func NewVehicleController(service service.VehicleService) *VehicleController {
	return &VehicleController{
		vehicleService: service,
	}
}

func (controller *VehicleController) Create(ctx *gin.Context) {
	createVehicleRequest := request.CreateVehicleRequest{}
	err := ctx.ShouldBindJSON(&createVehicleRequest)
	if err != nil {
		// Handle JSON binding error
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	vehicleID, err := controller.vehicleService.Create(createVehicleRequest)
	if err != nil {
		// Handle error from service layer
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data: map[string]int64{
			"vehicle_id": vehicleID,
		},
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
