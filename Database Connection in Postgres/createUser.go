package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(ctx *gin.Context) {
	var input CreateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u := User{}

	u.Username = input.Username
	u.Password = input.Password

	_, err := u.SaveUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":"User Created",
	})

}
