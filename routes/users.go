package routes

import (
	"Luc1808/goEvents/models"
	"Luc1808/goEvents/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Creating user failed.", "error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user."})
		return
	}

	err = user.VerifyCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "User is not authorized."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "User is not authorized."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})
}
