package request

type CreateDriverRequest struct {
	FirstName     string `json:"firstName" validate:"required"`
	LastName      string `json:"lastName"`
	Surname       string `json:"surname"`
	LicenseNumber string `json:"licenseNumber" validate:"required"`
	Email         string `json:"email"`
	Gender        string `json:"gender" validate:"required"`
	Dob           string `json:"dob" validate:"required"`
	MobileNo      string `json:"mobileNo" validate:"required"`
	Address       string `json:"address" validate:"required"`
	Language      string `json:"language" validate:"required"`
	AccountNumber string `json:"accountNumber" validate:"required"`
	DeviceID      string `json:"deviceId" validate:"required"`
	DeviceType    string `json:"deviceType" validate:"required"`
}
