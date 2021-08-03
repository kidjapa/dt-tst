package utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetDefaultErrorMessage Get the default return message for echo.Context
func GetDefaultErrorMessage(c echo.Context) error {
	return c.String(http.StatusNotFound, fmt.Sprintf("%d", 0))
}
