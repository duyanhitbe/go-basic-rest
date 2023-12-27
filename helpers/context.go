package helpers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIdFromParam retrieves the ID from the specified key in the given gin.Context.
// It parses the ID as an int64 from the context parameter with the specified key.
// If the parsing fails, it returns nil and sets the response status to http.StatusBadRequest.
// Otherwise, it returns a pointer to the parsed ID.
func GetIdFromParam(context *gin.Context, key string) *int64 {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	return &id
}
