package users

import (
	"github.com/gin-gonic/gin"
	"github.com/nictes1/Microservices-Go/bookstore_users-api/domain/users"
	"github.com/nictes1/Microservices-Go/bookstore_users-api/services"
	"github.com/nictes1/Microservices-Go/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func getUserId(userIdParameter string)(int64, *errors.RestErr){
	userID, userErr := strconv.ParseInt(userIdParameter, 10, 64)
	if userErr != nil {
		return 0,errors.NewBadRequestError("invalid user id")
	}
	return userID, nil
}

func Create(c *gin.Context)  {
	var user users.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil{
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context){
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil{
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context){
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil{
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil{
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "delete"})
}

func Search(c *gin.Context){
	status := c.Query("status")

	user, err := services.Search(status)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)
}