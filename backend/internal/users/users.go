package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	UserName string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password`
}

var users = []user {
	{UserName: "testuser1", Email: "test1@gamil.com", Password: "123"},
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)

}