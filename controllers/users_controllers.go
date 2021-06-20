package controllers

import (
	"bookstore_users_api/domain/users"
	"bookstore_users_api/services"
	errors "bookstore_users_api/utills/error"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetUser(c *gin.Context){
	user_id,UserErr:=strconv.ParseInt(c.Param("user_id"),10,64)
	if UserErr!= nil{
		err:=errors.NewBadRequestError("User id should be a number")
		c.JSON(err.Code,err)
		return
	}
	result,saveErr:= services.GetUser(user_id)
	if saveErr!=nil{
		c.JSON(saveErr.Code,saveErr)
		return
	}	
	c.JSON(http.StatusCreated,result)

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
		resterr:= errors.NewBadRequestError("invalid json body")
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