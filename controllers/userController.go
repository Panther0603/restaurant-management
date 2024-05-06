package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetUsers() gin.HandlerFunc {

	return func(ctx *gin.Context) {

	}
}

func GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func HashPassword(password string) string {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panicln(err.Error())
	}
	return string(hasedPassword)

}

func VerifyPassword(hashedPassword string, providedPassword string) (verified bool, msg string) {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))

	if err != nil {
		return false, err.Error()
	}
	return true, ""
}
