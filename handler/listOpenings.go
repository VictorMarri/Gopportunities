package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "List Openings",
	})
}
