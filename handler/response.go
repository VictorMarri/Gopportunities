package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendError(context *gin.Context, code int, message string) {
	context.Header("Content-Type", "application/json; charset=UTF-8")
	context.JSON(code, gin.H{"error": message})
}

func SendSuccess(context *gin.Context, operation string, data interface{}) {
	context.Header("Content-Type", "application/json; charset=UTF-8")
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully processed %s", operation),
		"data":    data,
	})
}
