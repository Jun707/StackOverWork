package server

import (
	"net/http"
	"stack/web/backend/internal/store"

	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err !=nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {"err":err.Error()})
		return
	}
	store.Users = append(store.Users, user)
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

	for _, u := range store.Users {
		if u.UserName == user.UserName && u.Password == user.Password {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Signed in success",
				"jwt": "123456789",
			})
			return
		}
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Sign in failed."})
}