package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func main() {
	server := echo.New()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	p := make([]User, 0)
	p = append(p, User{
		ID:        0,
		FirstName: "Harry",
		LastName:  "Potter",
		Age:       18,
	})
	p = append(p, User{
		ID:        2,
		FirstName: "Hermione",
		LastName:  "Granger",
		Age:       18,
	})

	// MIDDLEWARE
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// ROUTES
	server.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, p)
	})

	address := fmt.Sprintf("%s:%s", os.Getenv("IP"), os.Getenv("PORT"))

	server.Logger.Fatal(server.Start(address))
}
