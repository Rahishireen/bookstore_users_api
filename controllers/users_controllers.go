package controllers

import (
	"bookstore_users_api/domain/users"
	"bookstore_users_api/services"
	errors "bookstore_users_api/utills"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"implement me!!")

}

func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"implement me!!")

}

func CreateUser(c *gin.Context){
	var user users.User
	/*fmt.Println(user)
	bytes,err:= ioutil.ReadAll(c.Request.Body)
	if err!= nil{
		return
	}
	if err:=json.Unmarshal(bytes,&user);err!=nil{
		fmt.Println(err.Error())
	return

	}
	*/
	if err:= c.ShouldBindJSON(&user);err!=nil{
		resterr:= errors.RestErr{
			Message:"invalid json body",
			Code:http.StatusBadRequest,
			Error:"bad_request",
		}
		c.JSON(resterr.Code,resterr)
		return
	}
	result,saveErr:= services.CreateUser(user)
	if saveErr!=nil{
		c.JSON(saveErr.Code,saveErr)
		return
	}	
	c.JSON(http.StatusCreated,result)
}