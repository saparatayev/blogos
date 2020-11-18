package crud

import (
	"blogos/src/api/models"
	"blogos/src/api/utils/channels"

	"github.com/jinzhu/gorm"
)

type repositoryPostsCRUD struct {
	db *gorm.DB
}

func NewRepositoryPostsCRUD(db *gorm.DB) *repositoryPostsCRUD {
	return &repositoryPostsCRUD{
		db: db,
	}
}

func (r *repositoryPostsCRUD) Save(post models.Post) (models.Post, error) {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = r.db.Debug().Model(&models.Post{}).Create(&post).Error
		if err != nil {
			ch <- false

			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return post, nil
	}

	return models.Post{}, err
}
