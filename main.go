package main

import (
	"fmt"
	"os"
	"rahulProj/student/controllers"
	"rahulProj/student/resources/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	router := gin.Default()
	router.GET("/", controllers.StudentDetails)
	router.GET("/students/:id", controllers.StudentByID)
	router.POST("/register", controllers.RegisterStudent)
	router.GET("./unregister/:id", controllers.UnregisterStudent)

	godotenv.Load()

	fmt.Println("yaml files data")
	fmt.Println(os.Getenv("DB_NAME"))

	database.DatabaseInit()

	router.Run(":8080")
}
