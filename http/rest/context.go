package rest

import (
	"github.com/ingoaf/http-api/adding"
	"github.com/ingoaf/http-api/listing"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	Lister listing.Service
	Adder  adding.Service
}
