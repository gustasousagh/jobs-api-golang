package handler

import "github.com/gin-gonic/gin"

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, "Update")
}