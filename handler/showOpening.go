package handler

import (
	"fmt"
	"github.com/VictorMarri/Gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowOpeningHandler(context *gin.Context) {

	id := context.Query("id")

	if id == "" {
		SendError(context, http.StatusBadRequest,
			errParamIsRequired("id", "QueryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		SendError(context, http.StatusNotFound, fmt.Sprint("Failed to find register with id: ", id))
		return
	}

	SendSuccess(context, "Show opening", opening)
}
