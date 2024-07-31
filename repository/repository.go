package repository

import "tdd/models"

type Repository interface {
	Save(data *models.Data) error
}

type MemoryRepository struct {
	data []models.Data
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data: make([]models.Data, 0),
	}
}

func (r *MemoryRepository) Save(data *models.Data) error {
	r.data = append(r.data, *data)
	return nil
}
