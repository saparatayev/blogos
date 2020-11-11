package crud

import (
	"blogos/src/api/models"
	"blogos/src/api/utils/channels"

	"github.com/jinzhu/gorm"
)

type repositoryUsersCRUD struct {
	db *gorm.DB
}

func NewRepositoryUsersCRUD(db *gorm.DB) *repositoryUsersCRUD {
	return &repositoryUsersCRUD{
		db: db,
	}
}

func (r *repositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false

			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}

	return models.User{}, err
}

func (r *repositoryUsersCRUD) FindAll() ([]models.User, error) {
	var err error

	users := []models.User{}

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
		if err != nil {
			ch <- false

			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return users, nil
	}

	return nil, err
}
