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

func (r *repositoryPostsCRUD) FindAll() ([]models.Post, error) {
	var err error

	posts := []models.Post{}

	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = r.db.Debug().Model(&models.Post{}).Limit(100).Find(&posts).Error
		if err != nil {
			ch <- false

			return
		}

		if len(posts) > 0 {
			for i, _ := range posts {
				err = r.db.Debug().Model(&models.User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
				if err != nil {
					ch <- false

					return
				}
			}
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return posts, nil
	}

	return nil, err
}
