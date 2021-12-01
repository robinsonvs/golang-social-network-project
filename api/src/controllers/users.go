package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepositoryUsers(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprint("Inserted ID: %d", userID)))

}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find all users"))
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
