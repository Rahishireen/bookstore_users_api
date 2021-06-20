package users

import (
	"bookstore_users_api/utills/date"
	errors "bookstore_users_api/utills/error"
	"fmt"
)
var(
	UserDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr{
	result:= UserDB[user.Id]
	if result==nil{
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found",user.Id))
	}
	user.Id=result.Id
	user.FirstName = result.FirstName
	user.LastName =result.LastName
	user.Email=result.Email
	user.DateCreated=date.GetTimeString()
	return nil
}

func (user *User) Save() *errors.RestErr{
	current:= UserDB[user.Id]
	if current!=nil{
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist",user.Id))
	}
	UserDB[user.Id] = user
	return nil
}