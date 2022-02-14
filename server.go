package main

import (
	"github.com/ingoaf/http-api/adding"
	"github.com/ingoaf/http-api/http/rest"
	"github.com/ingoaf/http-api/listing"
	"github.com/ingoaf/http-api/memory"
	"github.com/labstack/echo/v4"
)

func main() {
	s := new(memory.Storage)

	adder := adding.NewService(s)
	lister := listing.NewService(s)

	e := echo.New()
	e.HideBanner = true
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &rest.Context{c, lister, adder}
			return next(cc)
		}
	})

	e.GET("/:name", rest.GetAttendee)
	e.POST("/attendee", rest.AddAttendee)

	e.Logger.Fatal(e.Start(":8080"))
}
