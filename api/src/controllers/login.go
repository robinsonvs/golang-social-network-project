package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/safety"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsers(db)
	userSavedInDatabase, err := repository.SearchByMail(user.Mail)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = safety.CheckPassword(userSavedInDatabase.Password, user.Password); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := authentication.CreateToken(userSavedInDatabase.ID)
	fmt.Println(token)
	w.Write([]byte(token))
}
