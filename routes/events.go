package routes

import (
	"net/http"
	"strconv"
	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func getEvents ( context *gin.Context ){
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Couldn't get events", "error": err.Error()})
		return
	}
	context.JSON(200, events)
}

func getEvent ( context *gin.Context ){
	id , err := strconv.ParseInt(context.Param("id"), 10 , 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Couldn't parse id", "error": err.Error()})
		return
	}
	
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Couldn't get events", "error": err.Error()})
		return
	}
	context.JSON(200, event)
}

func createEvent ( context *gin.Context ){
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Couldn't parese event", "error": err.Error()})
		return
	}
	event.UserID = context.GetInt64("userId")
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Couldn't create event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message" : "Event Created", "event": event})
}

func updateEvent ( context *gin.Context ){
	id , err := strconv.ParseInt(context.Param("id"), 10 , 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Couldn't parse id", "error": err.Error()})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Couldn't get events", "error": err.Error()})
		return
	}

	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Couldn't parese event", "error": err.Error()})
		return
	}
	event.ID = id
	err = event.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Couldn't Update Event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message" : "Event Updated", "event": event})
}


func deleteEvent ( context *gin.Context ){
	id , err := strconv.ParseInt(context.Param("id"), 10 , 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Couldn't parse id", "error": err.Error()})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Couldn't get event", "error": err.Error()})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Couldn't delete event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}
