package models

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Videogame struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string             `json:"name" bson:"name"`
	Description      string             `json:"description" bson:"descriptionr"`
	Platform         []string           `json:"platform" bson:"platform"`
	UserId           primitive.ObjectID `json:"user_id" bson:"user_id"`
}

type Videogames []Videogame

func (videogame *Videogame) DeleteVideogame() {
	err := mgm.Coll(videogame).Delete(videogame)
	if err != nil {
		fmt.Println(err.Error())
	}
}


func (videogame *Videogame) GetVideogame(id primitive.ObjectID) {
	coll := mgm.Coll(videogame)
	err := coll.FindByID(id, videogame)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (videogame *Videogame) GetVideoGamesByUser(userId primitive.ObjectID) Videogames {
	coll := mgm.Coll(videogame)
	videogames := Videogames{}
	err := coll.SimpleFind(&videogames, bson.D{{"user_id", userId}})
	if err != nil {
		fmt.Println(err.Error())
	}
	return videogames
}

func (videogame *Videogame) Save(videogameData Videogame) {
	if videogame.checkVideogameExists(videogameData.ID) == false {
		videogame.insert()
	} else {
		videogame.update(videogameData)
	}
}

func (videogame *Videogame) insert() {
	err := mgm.Coll(videogame).Create(videogame)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (videogame *Videogame) update(videogameData Videogame) {
	videogame.Name = videogameData.Name
	videogame.Description = videogameData.Description
	videogame.Platform = videogameData.Platform
	videogame.UserId = videogameData.UserId

	err := mgm.Coll(videogame).Update(videogame)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (videogame *Videogame) checkVideogameExists(id primitive.ObjectID) bool {
	coll := mgm.Coll(videogame)
	err := coll.FindByID(id, videogame)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		return true
	}
}
