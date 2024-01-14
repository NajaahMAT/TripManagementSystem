package service

import (
	"TripManagementSystem/data/request"
	"TripManagementSystem/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagRequest)
	Delete(tagId int)
	FindById(tagId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
