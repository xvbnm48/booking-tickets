package main

import (
	"booking-ticket/internal/controller"
	"booking-ticket/internal/repository"
	"booking-ticket/internal/service"
	"booking-ticket/logger"
	"booking-ticket/pkg/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("templates/*")
	dsn := "root:nYfw@-9*jf5P48@tcp(localhost:3306)/fr-ticket?parseTime=true"

	dbInstance, err := database.NewDatabase(dsn)
	if err != nil {
		// logger.Log.WithFields(logrus.Fields{
		// 	"error": err.Error(),
		// }).Error("Failed to connect to database")
		logger.Error("Failed to connect to database", err)
		return
	}
	defer dbInstance.GetConnection().Close()
	userRepo := repository.NewUserRepository(dbInstance.GetConnection())
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r.GET("/users", userController.GetAllUsers)
	r.GET("/user/:email", userController.GetUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.GET("/index", func(c *gin.Context) {
		users, err := userService.GetAllUsers()
		fmt.Println("isi users", users)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println("isi c", c)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"users": users,
		})
	})

	conn := dbInstance.GetConnection()
	if err := conn.Ping(); err != nil {
		// logger.Log.WithFields(logrus.Fields{
		// 	"error": err.Error(),
		// }).Error("Failed to ping to database")
		logger.Error("Failed to ping to database", err)
		return
	}
	r.Run(":6969")
}
