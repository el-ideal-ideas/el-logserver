package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)


// For check connection
func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "Success!")
}
