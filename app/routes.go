package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"time"
)

func APIRoutes(e *echo.Echo, db *sql.DB) {
	APIRoute := e.Group("/api")
	APIv1 := APIRoute.Group("/v1")

	APIv1.GET("", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		return c.JSON(200, map[string]string{
			"name":        "Children_app_api",
			"developer":   "Andika Nugraha",
			"version":     "v1.0",
			"status_code": fmt.Sprintf("%d", res.Status),
			"time":        time.Now().Format("2006/01/01 - 15:04:05"),
			"protocol":    req.Proto,
			"ip":          c.RealIP(),
			"method":      req.Method,
			"url":         fmt.Sprintf("%s", req.URL),
			"bytes_out":   fmt.Sprintf("%d", res.Size),
			"server_type": "Testing",
		})
	})
}
