package controller

import (
	"TripManagementSystem/data/request"
	"TripManagementSystem/data/response"
	"TripManagementSystem/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DVMappingController struct {
	dvMappingService service.DMMappingService
}

func NewDVMappingController(service service.DMMappingService) *DVMappingController {
	return &DVMappingController{
		dvMappingService: service,
	}
}

func (controller *DVMappingController) Create(ctx *gin.Context) {
	createDVMappingRequest := request.CreateDVMappingRequest{}
	err := ctx.ShouldBindJSON(&createDVMappingRequest)
	if err != nil {
		// Handle JSON binding error
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	mappingID, err := controller.dvMappingService.Create(createDVMappingRequest)
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
			"mapping_id": mappingID,
		},
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
