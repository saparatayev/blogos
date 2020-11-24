package repository

import "blogos/src/api/models"

type PostRepository interface {
	Save(models.Post) (models.Post, error)
	FindAll() ([]models.Post, error)
	FindById(uint64) (models.Post, error)
	Update(uint64, models.Post) (int64, error)
	Delete(uint64) (int64, error)
}
