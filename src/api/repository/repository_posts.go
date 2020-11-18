package repository

import "blogos/src/api/models"

type PostRepository interface {
	Save(models.Post) (models.Post, error)
	// FindAll() ([]models.Post, error)
	// FindById(uint32) (models.Post, error)
	// Update(uint32, models.Post) (int64, error)
	// Delete(uint32) (int64, error)
}
