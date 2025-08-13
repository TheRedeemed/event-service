package routes

import (
	"net/http"
	"strconv"
	"the-redeemed/event-service/models"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured while getting event"})
		return
	}

	// if event.UserId != userId {
	// 	context.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to register for this event"})
	// 	return
	// }

	err = event.RegisterForEvent(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured while registering for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered event"})
}

func unregisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	var event models.Event
	event.Id = eventId

	err = event.CancelRegisteration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured while cancelling event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Cancelled"})
}
