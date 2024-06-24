package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the events. Please try later"})

		return
	}
	context.JSON(http.StatusOK, events)
}
func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't fetch the id. Please try later"})

		return
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't fetch the event. Please try later"})

		return
	}
	context.JSON(http.StatusOK, event)

}
func createEvent(context *gin.Context) {
	userId := context.GetInt64("UserID")

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data"})

		return
	}
	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't create event. Please try later"})

		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the id. Please try later"})

		return
	}
	userId := context.GetInt64("UserID")

	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't fetch the event. Please try later"})

		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorize to update event."})

		return
	}
	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the event data. Please try later"})

		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't update the event data. Please try later"})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the id. Please try later"})

		return
	}
	userId := context.GetInt64("UserID")

	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't fetch the event. Please try later"})

		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorize to delete event."})

		return
	}
	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't delete the event. Please try later"})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
