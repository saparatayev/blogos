package auto

import "blogos/src/api/models"

var users = []models.User{
	models.User{Nickname: "John Doe", Email: "johndoe@mail.com", Password: "123456789"},
}

var posts = []models.Post{
	models.Post{
		Title:   "Test title 1",
		Content: "Hellow world content!",
	},
}
