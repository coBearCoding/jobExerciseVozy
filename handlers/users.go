package handlers

import (
	"encoding/json"
	"errors"
	"jobExercise/models"
	"jobExercise/validations"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendNotFound(w, err)
	} else {
		user.Save(user)
		models.SendData(w, user)
	}
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetUsers()
	models.SendData(w, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		models.SendNotFound(w, err)
	} else {
		validUserId := validations.CheckIfParameterEmptyOrNil(user.ID.Hex())
		if validUserId {
			user.GetUser(user.ID)
			models.SendData(w, user)
		}else{
			models.SendNotFound(w, errors.New("parameter id can't be empty"))
		}
	}
}

func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendNotFound(w, err)
	} else {
		validUserId := validations.CheckIfParameterEmptyOrNil(user.ID.Hex())
		if validUserId{
			user.Save(user)
			models.SendData(w, user)
		}else{
			models.SendNotFound(w, errors.New("parameter id can't be empty"))
		}

	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		models.SendNotFound(w, err)
	} else {
		validUserId := validations.CheckIfParameterEmptyOrNil(user.ID.Hex())
		if validUserId{
			user.DeleteUser()
			models.SendData(w, user)
		}else{
			models.SendNotFound(w, errors.New("parameter id can't be empty"))
		}

	}
}
