package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"jobExercise/models"
	"jobExercise/validations"
	"net/http"
)

func RegisterVideogame(w http.ResponseWriter, r *http.Request) {
	videogame := models.Videogame{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&videogame); err != nil {
		models.SendNotFound(w, err)
	} else {
		videogame.Save(videogame)
		models.SendData(w, videogame)
	}
}

func GetVideogame(w http.ResponseWriter, r *http.Request) {
	videogame := models.Videogame{}
	err := json.NewDecoder(r.Body).Decode(&videogame)
	if err != nil {
		models.SendNotFound(w, err)
	} else {
		videogame.GetVideogame(videogame.ID)
		models.SendData(w, videogame)
	}
}

func ListVideogames(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	videogame := models.Videogame{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		models.SendNotFound(w, err)
	} else {
		validUserId := validations.CheckIfParameterEmptyOrNil(videogame.UserId.Hex())
		if validUserId {
			userId := user.ID
			videogames := videogame.GetVideoGamesByUser(userId)
			models.SendData(w, videogames)
		} else {
			models.SendNotFound(w, errors.New("user_id can't be empty"))
		}

	}
}

func UpdateVideogameInfo(w http.ResponseWriter, r *http.Request) {
	videogame := models.Videogame{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&videogame); err != nil {
		models.SendNotFound(w, err)
	} else {
		validUserId := validations.CheckIfParameterEmptyOrNil(videogame.UserId.Hex())
		validVideogameId := validations.CheckIfParameterEmptyOrNil(videogame.ID.Hex())

		if validUserId && validVideogameId {
			fmt.Println(videogame.UserId.Hex())
			fmt.Println(videogame.ID.Hex())
			fmt.Println(validUserId)
			fmt.Println(validVideogameId)
			videogame.Save(videogame)
			models.SendData(w, videogame)
		} else {
			fmt.Println(validUserId)
			fmt.Println(validVideogameId)
			models.SendNotFound(w, errors.New("id and user_id can't be empty"))
		}

	}
}

func DeleteVideogame(w http.ResponseWriter, r *http.Request) {
	videogame := models.Videogame{}
	err := json.NewDecoder(r.Body).Decode(&videogame)
	if err != nil {
		models.SendNotFound(w, err)
	} else {
		validUserId := validations.CheckIfParameterEmptyOrNil(videogame.ID.Hex())
		if validUserId {
			videogame.DeleteVideogame()
			models.SendData(w, videogame)
		} else {
			models.SendNotFound(w, errors.New("parameter id can't be empty"))
		}

	}
}
