package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustasousagh/jobs-api-golang/schema"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	opening := schema.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("error creating opening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating opening database")
		return
	}
	sendSuccess(ctx, "create-opening", opening)
}
