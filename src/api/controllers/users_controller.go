package controllers

import (
	"blogos/src/api/database"
	"blogos/src/api/models"
	"blogos/src/api/repository"
	"blogos/src/api/repository/crud"
	"blogos/src/api/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		users, err := usersRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		responses.JSON(w, http.StatusOK, users)
	}(repo)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		user, err = usersRepository.Save(user)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))

		responses.JSON(w, http.StatusCreated, user)
	}(repo)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("An user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}