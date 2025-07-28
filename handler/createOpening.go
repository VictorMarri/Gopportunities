package handler

import (
	"github.com/VictorMarri/Gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

	//Pointer to request data structure
	if err := context.BindJSON(&request); err != nil {
		logger.Errorf("Failed to bind request: %v", err)
		SendError(context, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("CreateOpeningHandler Validation Error: %v", err)
		SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Failed to create opening request: %v", err)
		SendError(context, http.StatusInternalServerError, "Failed to save data on database")
	}

	SendSuccess(context, "Created", opening)

}
