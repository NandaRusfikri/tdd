package usecase

import (
	"tdd/models"
	"tdd/repository"
)

type UseCase interface {
	CreateData(data *models.Data) error
}

type useCase struct {
	repo repository.Repository
}

func NewUseCase(repo repository.Repository) UseCase {
	return &useCase{repo: repo}
}

func (uc *useCase) CreateData(data *models.Data) error {
	return uc.repo.Save(data)
}
