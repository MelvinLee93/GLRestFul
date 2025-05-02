package jwtauth

import (
	"net/http"
    "time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Jwtcustomclaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func Login (c echo.Context) error{
	username := c.FormValue("username")
    password := c.FormValue("password")
    if username != "Caelus" || password != "baseball" {
		return echo.ErrUnauthorized
	}

    claims := &Jwtcustomclaims{
		"Galactic Baseballer",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72).UTC()),
		},
	}
    
	//JWT Token creation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("penacony"))
	if err != nil {
		return err
	}

    return c.JSON(200, map[string]string{
        "token": t,
    })
}

func Verboten(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(*Jwtcustomclaims)
    name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome, Guest!")
}