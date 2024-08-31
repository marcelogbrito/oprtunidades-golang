package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcelogbrito/oprtunidades-golang/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	// binding to populate
	ctx.BindJSON(&request)
	//validate
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	//logger.Infof("request received: %+v", request)

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
	//creating in database
	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating openinh: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening")
		return
	}

	sendSuccess(ctx, "create opening", opening)
}
