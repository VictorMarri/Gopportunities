package handler

import (
	"fmt"
	"github.com/VictorMarri/Gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteOpeningHandler(context *gin.Context) {
	id := context.Query("id")

	if id == "" {
		SendError(context, http.StatusBadRequest,
			errParamIsRequired("id", "QueryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	//Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("Failed to find id: %v", err)
		SendError(context, http.StatusNotFound, fmt.Sprint("Failed to find id: ", id))
		return
	}
	//Delete opening
	if err := db.Delete(&opening).Error; err != nil {
		logger.Errorf("Failed to delete register with id: %v", err)
		SendError(context, http.StatusInternalServerError, fmt.Sprintf("Failed to delete register with id: %v", err))
		return
	}

	SendSuccess(context, "Delete", opening)
}
