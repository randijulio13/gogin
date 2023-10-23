package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/randijulio13/gogin/exception"
	"github.com/randijulio13/gogin/model"
	"github.com/randijulio13/gogin/request"
)

func Login(c *gin.Context) {
	var request request.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	var user model.User
	if query := user.GetByNip(request.Nip); query.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "nip not found"})
		return
	}

	if err := user.ComparePassword(request.Password); err != nil {
		panic(exception.ValidationException{Error: "wrong password"})
	}

	token, err := user.GenerateToken()
	if err != nil {
		fmt.Println(err.Error())
		panic("failed generate token.")
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
