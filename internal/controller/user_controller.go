package controller

import (
	"booking-ticket/internal/model"
	"booking-ticket/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUsers(ctx *gin.Context)
	GetUserByEmail(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{service: service}
}

func (c *userController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *userController) GetUserByEmail(ctx *gin.Context) {
	// var req model.UserRequestByEmail
	// if err := ctx.BindJSON(req.Email); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	email := ctx.Param("email")
	fmt.Println("isi param email", email)
	user, err := c.service.GetUserByEmail(email)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *userController) CreateUser(ctx *gin.Context) {
	var req model.CreateUserRequest
	fmt.Println("isi req", req)
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("isi req", req)
	status, err := c.service.CreateUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Success Crete user", "Status": status})
}

func (c *userController) GetUserById(ctx *gin.Context) {
	// id from query parameter
	// id := ctx.Param("id")
	id := ctx.Param("id")
	fmt.Println("isi param id", id)
	user, err := c.service.GetUserById(id)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, model.UserResponseStruct{
		Status:  "Success Get User",
		Message: "User found",
		Data:    user,
	})
}
