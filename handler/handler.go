package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, "Create")
}
func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, "Show")
}
func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, "Delete")
}
func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, "Update")
}
func ListOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, "List")
}
