package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *Service) createUser(c echo.Context) error {
	m := new(Member)

	if err := c.Bind(m); err != nil {
		return err
	}

	err := s.Store.createUser(m.Name, m.Phone, m.Email)
	if err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, m)
}
