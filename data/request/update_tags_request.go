package request

type UpdateTagRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required" json:"name"`
}
