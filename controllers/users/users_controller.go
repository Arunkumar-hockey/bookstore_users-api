package users

import (
	"fmt"

	"net/http"
	"userapi/domain/users"
	"userapi/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		//TODO: return bad request to the caller.
		return
	}

	result, saveErr := services.CreateUser(user)
	if  saveErr != nil {
		//TODO: handle user creation error
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
