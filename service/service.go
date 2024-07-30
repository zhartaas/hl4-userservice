package service

import (
	"hl4-user_service/repository"
	"hl4-user_service/userdomain"
)

type Service struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetAllUsers() (users []userdomain.Entity, err error) {
	users, err = s.repository.GetAllUsers()
	return
}
