package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gustasousagh/jobs-api-golang/schema"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(ctx *gin.Context) {
	body := Register{}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to generate hash",
		})
	}
	user := schema.User{
		Email:    body.Email,
		Password: string(hash),
	}
	result := db.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func Login(ctx *gin.Context) {
	body := Register{}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Failed to read body",
		})
		return
	}
	user := schema.User{}

	db.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Email ou password"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Email ou password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("bjkbjbbhibib78"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Ereror creating token"})
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func Validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
