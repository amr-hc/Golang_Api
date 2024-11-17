package routes

import (
	"net/http"
	"example.com/api/models"
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