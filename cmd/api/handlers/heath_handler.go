package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HealthCheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK!")
}
