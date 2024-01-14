package request

type CreateVehicleRequest struct {
	RegistrationNumber string `json:"RegistrationNumber" validate:"required"`
	Type               string `json:"type" validate:"required"`
	ModelMake          string `json:"ModelMake" validate:"required"`
	BrandModel         string `json:"BrandModel" validate:"required"`
	Color              string `json:"Color" validate:"required"`
	Year               int    `json:"Year" validate:"required"`
	SeatingCapacity    int    `json:"SeatingCapacity" validate:"required"`
	EngineCapacity     int    `json:"EngineCapacity" validate:"required"`
}
