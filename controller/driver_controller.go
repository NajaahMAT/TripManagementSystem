package controller

import (
	"TripManagementSystem/data/request"
	"TripManagementSystem/data/response"
	"TripManagementSystem/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DriverController struct {
	driverService service.DriverService
}

func NewDriverController(service service.DriverService) *DriverController {
	return &DriverController{
		driverService: service,
	}
}

func (controller *DriverController) Create(ctx *gin.Context) {
	createDriverRequest := request.CreateDriverRequest{}
	err := ctx.ShouldBindJSON(&createDriverRequest)
	if err != nil {
		// Handle JSON binding error
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	driverID, err := controller.driverService.Create(createDriverRequest)
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
			"driver_id": driverID,
		},
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
