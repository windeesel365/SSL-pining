package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SecuredHandler(c echo.Context) error {
	return c.String(http.StatusOK, "This is a valid server.")
}
