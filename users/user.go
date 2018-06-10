package users

import (
	"time"

	"github.com/VitorLuizC/ComuniCode-Hackathon/server/db"

	"github.com/mitchellh/mapstructure"
)

var UserCollection = db.RepositoryDefMap{
	"name":    "users",
	"hashKey": "_id",
}

type User struct {
	Id       string    `json:"-" bson:"_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	UpdateAt time.Time `json:"updateAt"`
}

func GetUser(id string) (*User, error) {
	var user *User
	res, err := db.MongoRepoBuilder(UserCollection).FindOne(id)

	if err != nil {
		return nil, err
	}
	if err = mapstructure.Decode(res, &user); err != nil {
		return nil, err
	}
	return user, nil
}
