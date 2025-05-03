package main

import (
	"Prototype/actions"
	"Prototype/auth"
	"Prototype/config"
	_ "log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e.Use(middleware.Logger())
    e.Use(middleware.Recover())

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	GuestCRUD := e.Group("/")
	{
		GuestCRUD.GET("", jwtauth.Accessible)
		GuestCRUD.GET(":id", actions.FindEmployee)
	}

	e.POST("/login", jwtauth.Login)
    CRUDAction := e.Group("/restricted")
	{
		config := echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(jwtauth.Jwtcustomclaims)
			},
			SigningKey: []byte("penacony"),
		}
		CRUDAction.Use(echojwt.WithConfig(config))
		CRUDAction.POST("/", actions.AddEmployee)
		CRUDAction.GET("", jwtauth.Verboten)
		CRUDAction.GET("/:id", actions.FindEmployee)
		CRUDAction.PATCH("/:id", actions.UpdateEmployee)
		CRUDAction.DELETE("/:id", actions.DeleteEmployee)
	}
    e.Logger.Fatal(e.Start(":8080"))
}