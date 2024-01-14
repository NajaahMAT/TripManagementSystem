package request

type CreateTagsRequest struct {
	Name string `validate:"required" json:"name"`
}
