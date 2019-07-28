package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func createUser(c echo.Context) error {
	m := new(Member)

	if err := c.Bind(m); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, m)
}
