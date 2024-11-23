package routes

import (
	"net/http"

	"example.com/api/models"
	"example.com/api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't parese event", "error": err.Error()})
		return
	}
	err = user.Signup()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't create event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user Created", "user": user})
}


func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't parese event", "error": err.Error()})
		return
	}
	err = user.Login()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "invalid password"})
		return
	}

	token , err := utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login Successful", "token": token})
}