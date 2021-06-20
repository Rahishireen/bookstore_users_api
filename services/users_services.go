package services

import (
	"bookstore_users_api/domain/users"
	errors "bookstore_users_api/utills/error"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr){

	if err:=user.Validate();err!=nil{
		return nil,err
	}
	if err:=user.Save();err!=nil{
		return nil,err
	}

	return &user,nil
}

func GetUser(user_id int64) (*users.User, *errors.RestErr){
	result:= &users.User{Id:user_id	}
	if err:=result.Get();err!=nil{
		return nil,err
	}
	return result,nil
	
}