package users

import (
	"errors"
	"log"
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
	Password string    `json:"password"`
	Cpf      string    `json:"cpf"`
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

func FindByEmail(email string) (*User, error) {
	var data *User
	res, err := db.MongoRepoBuilder(UserCollection).FindByQuery("email", email)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("user not found")
	}
	if err := mapstructure.Decode(res, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (user *User) Save() error {
	log.Printf("save user: %v", user)
	if err := db.MongoRepoBuilder(UserCollection).Save(user); err != nil {
		return err
	}
	return nil
}
