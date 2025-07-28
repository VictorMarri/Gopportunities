package handler

import (
	"github.com/VictorMarri/Gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningHandler(context *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		SendError(context, http.StatusInternalServerError, "Error listing openings")
		return
	}

	SendSuccess(context, "List openings", openings)
}
