package models

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
	Admin            bool   `json:"admin" bson:"admin"`
}

type Users []User

func GetUsers() Users {
	user := &User{}
	coll := mgm.Coll(user)
	users := Users{}
	err := coll.SimpleFind(&users, bson.D{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return users
}

func (user *User) GetUser(id primitive.ObjectID) {
	coll := mgm.Coll(user)
	err := coll.FindByID(id, user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (user *User) DeleteUser() {
	err := mgm.Coll(user).Delete(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (user *User) Save(userData User) {
	if user.checkUserExists(user.ID) == false {
		user.insert()
	} else {
		user.update(userData)
	}
}

func (user *User) update(userData User) {
	user.Username = userData.Username
	user.Email = userData.Email
	user.Password = userData.Password
	user.Admin = userData.Admin

	err := mgm.Coll(user).Update(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (user *User) insert() {
	err := mgm.Coll(user).Create(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (user *User) checkUserExists(id primitive.ObjectID) bool {
	coll := mgm.Coll(user)
	err := coll.FindByID(id, user)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		return true
	}
}
