package GinUserService

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func marshalJson(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return nil, err
	}
	return data, nil
}

func GinResponseHandler(ctx *gin.Context, data interface{}, err interface{}, statusCode int, messages ...string) {
	ctx.Header("Content-Type", "application/json; charset=UTF-8")
	ctx.Status(statusCode)

	message := ""
	for _, m := range messages {
		if message == "" {
			message = m
		} else {
			message = message + " - " + m
		}
	}

	response := struct {
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
		Message string      `json:"message,omitempty"`
	}{
		Data:    data,
		Error:   err,
		Message: message,
	}

	serialized, jsonErr := marshalJson(response)
	if jsonErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Serialization Error"})
		return
	}

	ctx.Data(statusCode, "application/json; charset=UTF-8", serialized)
}
