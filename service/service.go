package service

import (
	modelApi "github.com/kaburov38/GoShort/model/api"
	"github.com/kaburov38/GoShort/model"
	repository "github.com/kaburov38/GoShort/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Insert(request modelApi.InsertRequest) error {
	return s.repo.Insert(model.Mapping{Source: request.Source, Target: request.Target})
}

func (s *Service) Find(request modelApi.FindRequest) (response modelApi.MappingResponse, err error) {
	entity, err := s.repo.Find(request.Source)
	if err != nil{
		return 
	}

	response = modelApi.MappingResponse{Source: entity.Source, Target: entity.Target}
	return
}

func (s *Service) Update(request modelApi.UpdateRequest) error {
	return s.repo.Update(model.Mapping{Source: request.Source}, model.Mapping{Source: request.Source, Target: request.Target})
}

func (s *Service) Delete(request modelApi.DeleteRequest) error {
	entity, err := s.repo.Find(request.Source)
	if err != nil{
		return err
	}

	return s.repo.Delete(entity)
}