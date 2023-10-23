package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/randijulio13/gogin/database"
	"github.com/randijulio13/gogin/model"
	"github.com/randijulio13/gogin/request"
)

type EditRequest struct {
	Password    string   `json:"password,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

func CreateUser(c *gin.Context) {
	var request request.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	if err := request.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	u := model.User{
		Nip:      request.Nip,
		Position: request.Position,
		Name:     request.Name,
		Email:    request.Email,
	}

	u.SetPassword(request.Password)
	u.SetPermissions(request.Permissions)
	user, err := u.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created", "data": user})
}

func EditUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	query := user.GetById(id)

	if query.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	var request EditRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	if len(request.Permissions) != 0 {
		var permissions []model.Permission
		model.GetPermissions(&permissions, request.Permissions)
		user.Permissions = permissions
	}

	if len(request.Password) != 0 {
		user.SetPassword(request.Password)
	}

	user.UpdateUser()
	c.JSON(http.StatusOK, gin.H{"message": "user data updated", "user": user})
}

func GetAllUser(c *gin.Context) {
	db := database.DB
	var user []model.User
	db.Preload("Permissions").Find(&user)
	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": user})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	query := user.GetById(id)

	if query.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": user})
}
