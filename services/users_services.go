package services

import (
	"bookstore_users_api/domain/users"
	errors "bookstore_users_api/utills"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr){

	return &user,nil
}