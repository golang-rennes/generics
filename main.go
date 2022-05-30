package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Person struct {
	Fistname string `json:"first_name"`
	Lastname string `json:"last_name"`
}
type Cat struct {
	Name string `json:"name"`
}

type Handler[T any] struct {
	entities []T
}

func (h *Handler[T]) fetch(c echo.Context) error {
	return c.JSON(http.StatusOK, h.entities)
}

func main() {
	e := echo.New()

	personHandler := Handler[Person]{[]Person{
		{Fistname: "Valentin", Lastname: "Deleplace"}, {Fistname: "Gwendal", Lastname: "Leclerc"}, {Fistname: "Marc", Lastname: "Guenneguez"},
	}}
	catHandler := Handler[Cat]{[]Cat{
		{Name: "Linux"}, {Name: "Houston"},
	}}

	e.GET("/persons", personHandler.fetch)
	e.GET("/cats", catHandler.fetch)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
