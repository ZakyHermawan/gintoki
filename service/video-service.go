package service

import (
	"gintoki/entity"
	"gintoki/repository"
)

type VideoService interface {
	Save(entity.Video) error
	Update(entity.Video) error
	Delete(entity.Video) error
	FindAll() ([]entity.Video, error)
}

type videoService struct {
	repository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		repository: repo,
	}
}

func (service *videoService) Save(video entity.Video) error {
	return service.repository.Save(video)
}

func (service *videoService) Update(video entity.Video) error {
	return service.repository.Update(video)
}

func (service *videoService) Delete(video entity.Video) error {
	return service.repository.Delete(video)
}

func (service *videoService) FindAll() ([]entity.Video, error) {
	return service.repository.FindAll()
}
