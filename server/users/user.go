package users

import (
	"errors"
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
	res, err := db.MongoRepoBuilder(UserCollection).FindById(id)

	if err != nil {
		return nil, err
	}
	if err = mapstructure.Decode(res, &user); err != nil {
		return nil, err
	}
	return user, nil
}

func FindByEmail(email string) (*Users, error) {
	res, err := db.MongoRepoBuilder(UserCollection).FindByQuery("email", email)
	if err != nil {
		return nil, nil
	}
	//Todo: convert res to user
	return nil, nil
}

func (user *User) Save() error {
	res, err := FindByEmail(user.Email)
	if err != nil {
		return err
	}
	if res != nil {
		return errors.New("email already registered")
	}
	if err = db.MongoRepoBuilder(UserCollection).Save(user); err != nil {
		return err
	}

	return nil
}
