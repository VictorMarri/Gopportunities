package handler

import (
	"github.com/VictorMarri/Gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOpeningHandler(context *gin.Context) {
	request := UpdateOpeningRequest{}

	context.Bind(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	id := context.Query("id")

	if id == "" {
		SendError(context, http.StatusBadRequest, errParamIsRequired("id", "QueryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		SendError(context, http.StatusNotFound, "Opening not found")
		return
	}

	//Update opening

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary != 0 {
		opening.Salary = request.Salary
	}

	//Save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error saving opening: %v", err.Error())
		SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccess(context, "Update opening", id)
}
