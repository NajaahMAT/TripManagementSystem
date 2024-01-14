package service

import (
	"TripManagementSystem/data/request"
	"TripManagementSystem/data/response"
	"TripManagementSystem/helper"
	"TripManagementSystem/model"
	"TripManagementSystem/repository"

	validator "github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}

	t.TagsRepository.Save(tagModel)
}

func (t *TagsServiceImpl) Update(tags request.UpdateTagRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}

func (t *TagsServiceImpl) Delete(tagId int) {
	t.TagsRepository.Delete(tagId)
}

func (t *TagsServiceImpl) FindById(tagId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}

	return tagResponse
}

func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}

		tags = append(tags, tag)
	}

	return tags
}
