package models

import (
	"errors"

	"github.com/joshuakwan/time-tracker/dal"
	"github.com/joshuakwan/time-tracker/utils"
	"gopkg.in/mgo.v2/bson"
)

const (
	userTableName = "Users"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetUsersByEmail queries user by an email address
func GetUsersByEmail(email string) ([]User, error) {
	var result []User
	err := db.GetSession().FindObjects(userTableName, "email:"+email, &result)
	return result, err
}

// GetAllUsers returns all users
func GetAllUsers() ([]User, error) {
	var result []User
	err := db.GetSession().GetAllObjects(userTableName, &result)
	return result, err
}

// CreateUser creates a new user
func CreateUser(name, email string) (*User, error) {
	existingUsers, err := GetUsersByEmail(email)

	if err != nil {
		return nil, err
	}

	if len(existingUsers) > 0 {
		return nil, errors.New("user " + email + " exists")
	}

	user := User{ID: utils.GenerateUUID(), Name: name, Email: email}
	err = db.GetSession().CreateObject(userTableName, user)
	return &user, err
}

// Delete removes a user itself from DB
func (user *User) Delete() error {
	_, err := GetUsersByEmail(user.Email)
	if err != nil {
		return err
	}

	err = db.GetSession().RemoveObject(userTableName, "email:"+user.Email)
	return err
}

// Update updates a user itself
func (user *User) Update() error {
	err := db.GetSession().UpdateObject(userTableName, bson.M{"email": user.Email}, user)
	return err
}
