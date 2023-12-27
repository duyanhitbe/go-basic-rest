package controllers

import (
	"net/http"

	"github.com/duyanhitbe/go-basic-rest/helpers"
	"github.com/duyanhitbe/go-basic-rest/models"
	"github.com/gin-gonic/gin"
)

// GetAllEvent retrieves all events from the database and returns them as a JSON response.
// It uses the GetAllEvents function from the models package to fetch the events.
// If an error occurs during the retrieval process, it returns an internal server error with the error message.
// Otherwise, it returns a status OK response with the events.
func GetAllEvent(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, events)
}

// CreateEvent is a handler function that creates a new event.
// It binds the JSON data from the request body to the event model,
// then calls the Create method on the event to persist it in the database.
// If any error occurs during the binding or creation process,
// it returns an appropriate JSON response with the error message.
// If the event is created successfully, it returns a JSON response with the created event.
func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = event.Create()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, event)
}

// GetOneEvent retrieves a single event by its ID from the database and returns it as JSON.
// If the ID parameter is not a valid integer, it returns a bad request error.
// If there is an error retrieving the event from the database, it returns an internal server error.
// Otherwise, it returns the event as JSON with a status code of 200 OK.
func GetOneEvent(context *gin.Context) {
	id := helpers.GetIdFromParam(context, "id")
	event, err := models.GetOneEventById(*id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

// UpdateOneEvent updates an event by its ID.
// It retrieves the event from the database using the provided ID,
// binds the JSON data from the request body to the event struct,
// and updates the event in the database.
// If any error occurs during the process, it returns an appropriate JSON response.
// Otherwise, it returns the updated event as a JSON response with status code 200.
func UpdateOneEvent(context *gin.Context) {
	id := helpers.GetIdFromParam(context, "id")
	event, err := models.GetOneEventById(*id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err = models.UpdateEventById(*id, event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func DeleteOneEventById(context *gin.Context) {
	id := helpers.GetIdFromParam(context, "id")
	event, err := models.DeleteEventById(*id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}
