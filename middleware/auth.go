package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gustasousagh/jobs-api-golang/config"
	"github.com/gustasousagh/jobs-api-golang/schema"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandler() {
	db = config.GetSqlite()
}

func Auth(ctx *gin.Context) {

	tokenString, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unaexpeted method %v", t.Header)
		}
		return []byte("bjkbjbbhibib78"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		user := schema.User{}

		db.First(&user, claims["sub"])
		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		ctx.Set("user", user)
		ctx.Next()

	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
