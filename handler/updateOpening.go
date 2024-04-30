package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustasousagh/jobs-api-golang/schema"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateRequest{}

	ctx.BindJSON(&request)

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id")
		return
	}
	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Salary <= 0 {
		opening.Salary = request.Salary
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if err := db.Save(&opening).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	sendSuccess(ctx, "update-opening", opening)
}
