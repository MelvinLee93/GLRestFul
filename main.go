package main

import (
 _ "log"
 "github.com/labstack/echo/v4"
 "Prototype/config"
 "Prototype/actions"
 _ "github.com/lib/pq"
)

type Response struct {
    Message string
    Status  bool
}

func main() {
    //Connecting to database
    e := echo.New()
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

    CRUDAction := e.Group("/users")
	CRUDAction.POST("/", actions.AddEmployee)
	CRUDAction.GET("/:id", actions.FindEmployee)
	CRUDAction.PATCH("/:id", actions.UpdateEmployee)
	CRUDAction.DELETE("/:id", actions.DeleteEmployee)

    e.Logger.Fatal(e.Start(":8080"))
}