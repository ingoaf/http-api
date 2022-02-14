package rest

import (
	"net/http"

	"github.com/ingoaf/http-api/adding"
	"github.com/ingoaf/http-api/listing"
	"github.com/labstack/echo/v4"
)

func GetAttendee(c echo.Context) error {
	cc := c.(*Context)
	givenName := cc.Param("name")

	if givenName == "" {
		return cc.JSON(http.StatusBadRequest, "no attendee name provided")
	}

	attendee := listing.Attendee{
		Name: givenName,
	}

	attendee, err := cc.Lister.GetAttendee(attendee.Name)
	if err != nil {
		return cc.JSON(http.StatusInternalServerError, err.Error())
	}

	return cc.JSON(http.StatusOK, attendee)
}

func GetAllAttendees(c echo.Context) error {
	cc := c.(*Context)

	attendees := cc.Lister.GetAttendees()

	return cc.JSON(http.StatusOK, attendees)
}

func AddAttendee(c echo.Context) error {
	cc := c.(*Context)

	attendee := adding.Attendee{}
	if err := cc.Bind(&attendee); err != nil {
		return cc.JSON(http.StatusBadRequest, nil)
	}

	err := cc.Adder.AddAttendee(attendee)
	if err != nil {
		return cc.JSON(http.StatusInternalServerError, err.Error())
	}

	return cc.JSON(http.StatusOK, "attendee created")
}
