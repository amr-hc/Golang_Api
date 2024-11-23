package routes

import (
	"net/http"
	"strconv"
	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId , err := strconv.ParseInt(context.Param("id"), 10 , 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Couldn't parse id", "error": err.Error()})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message" : "couldn't get event"})
        return
    }
	err = event.Register(userId)
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message" : "couldn't register event"})
        return
    }
	context.JSON(http.StatusOK, gin.H{"message": "Registered successfully"})
}
func cancelRegistration(context *gin.Context)  {
	userId := context.GetInt64("userId")
	eventId , err := strconv.ParseInt(context.Param("id"), 10 , 64)
	if err!= nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message" : "couldn't parse event ID"})
        return
    }
	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err!= nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message" : "couldn't cancel registration"})
        return
    }
	context.JSON(http.StatusOK, gin.H{"message": "Cancelled registration successfully"})

}