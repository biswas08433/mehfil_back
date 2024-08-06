package main

import "github.com/labstack/echo/v4"

func CustomH1(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("CH1", "blahblah")
		return next(c)
	}
}
