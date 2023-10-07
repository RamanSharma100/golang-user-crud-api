package controllers

import (
	"crudapi/pkgs/models"
	"crudapi/pkgs/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)

	user := CreateUser.CreateUser()
	res, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	Id, err := strconv.ParseInt(key, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	user, _ := models.GetUserById(Id)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	Id, err := strconv.ParseInt(key, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	UpdateUser := &models.User{}
	utils.ParseBody(r, UpdateUser)

	user, db := models.GetUserById(Id)

	if UpdateUser.Name != "" {
		user.Name = UpdateUser.Name
	}
	if UpdateUser.Email != "" {
		user.Email = UpdateUser.Email
	}
	if UpdateUser.Age != 0 {
		user.Age = UpdateUser.Age
	}
	if UpdateUser.Phone != "" {
		user.Phone = UpdateUser.Phone
	}
	if UpdateUser.Country != "" {
		user.Country = UpdateUser.Country
	}

	db.Save(&user)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	Id, err := strconv.ParseInt(key, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	user := models.DeleteUser(Id)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
