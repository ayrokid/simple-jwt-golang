package main

import (
	"RestFullApi/configs"
	"RestFullApi/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := configs.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/login", inDB.LoginHandler)
	router.GET("/person/:id", inDB.GetPerson)
	router.GET("/persons", inDB.GetPersons)
	router.POST("/person", inDB.CreatePerson)
	router.PUT("/person", inDB.UpdatePerson)
	router.DELETE("/person/:id", inDB.DeletePerson)
	router.Run(":4000")
}
