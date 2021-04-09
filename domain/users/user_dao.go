package users

import (
	"fmt"
	"github.com/nictes1/Microservices-Go/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreate = result.DateCreate
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email{
			return errors.NewBadRequestError(fmt.Sprintf("email %s alredy register", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d alredy exists", user.Id))
	}
	userDB[user.Id] = user
	return nil
}
