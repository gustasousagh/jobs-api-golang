package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustasousagh/jobs-api-golang/schema"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schema.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error Listing opening")
		return
	}

	sendSuccess(ctx, "list-opening", openings)
}
