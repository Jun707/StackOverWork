package server

import (
	"net/http"
	"stack/web/backend/internal/store"

	"github.com/gin-gonic/gin"
)

func retrieveUser(ctx *gin.Context) {
	users, err := store.GetUsers()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, users)
}

func signUp(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err !=nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	}

	if err := store.AddUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed up success.",
		"jwt": "123456789",
	})
}

func signIn(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err !=nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err":err.Error()})
		return
	}

	user, err := store.Authenticate(user.Email, user.Password)
	if err != nil {
	  ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sign in failed."})
	  return
	}
   
	ctx.JSON(http.StatusOK, gin.H{
	  "msg": "Signed in successfully.",
	  "jwt": "123456789",
	})
}