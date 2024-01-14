package controller

import (
	"TripManagementSystem/data/request"
	"TripManagementSystem/data/response"
	"TripManagementSystem/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TripController struct {
	tripService service.TripService
}

func NewTripController(service service.TripService) *TripController {
	return &TripController{
		tripService: service,
	}
}

func (controller *TripController) Create(ctx *gin.Context) {
	createTripRequest := request.CreateTripRequest{}
	err := ctx.ShouldBindJSON(&createTripRequest)
	if err != nil {
		// Handle JSON binding error
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	tripID, err := controller.tripService.Create(createTripRequest)
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
			"trip_id": tripID,
		},
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TripController) UpdateTripStatus(ctx *gin.Context) {
	updateTripsRequest := request.UpdateTripStatusRequest{}
	err := ctx.ShouldBindJSON(&updateTripsRequest)
	if err != nil {
		// Handle JSON binding error
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	tripId := ctx.Param("tripId")
	id, err := strconv.Atoi(tripId)
	if err != nil {
		// Handle error from service layer
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}
	updateTripsRequest.TripID = int64(id)

	err = controller.tripService.UpdateTravelStatus(updateTripsRequest)
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
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
