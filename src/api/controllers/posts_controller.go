package controllers

import (
	"blogos/src/api/auth"
	"blogos/src/api/database"
	"blogos/src/api/models"
	"blogos/src/api/repository"
	"blogos/src/api/repository/crud"
	"blogos/src/api/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postsRepository repository.PostRepository) {
		posts, err := postsRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		responses.JSON(w, http.StatusOK, posts)
	}(repo)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.Prepare()

	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postsRepository repository.PostRepository) {
		post, err = postsRepository.Save(post)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, post.ID))

		responses.JSON(w, http.StatusCreated, post)
	}(repo)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postsRepository repository.PostRepository) {
		post, err := postsRepository.FindById(pid)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, post)
	}(repo)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	post := models.Post{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.Prepare()

	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postsRepository repository.PostRepository) {
		rows, err := postsRepository.Update(pid, post)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, rows)
	}(repo)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postsRepository repository.PostRepository) {
		_, err := postsRepository.Delete(pid, uid)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		w.Header().Set("Entity", fmt.Sprintf("%d", pid))

		responses.JSON(w, http.StatusNoContent, "")
	}(repo)
}
