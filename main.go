package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Group
	g1 := e.Group("/g1")
	{
		g1.GET("/test", func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		})
	}

	// body json
	e.POST("/r1", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, u)
		}
		return c.JSON(http.StatusOK, u)
	})

	// param
	e.GET("/Param/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, name)
	})

	// QueryParam
	e.GET("/QueryParam", func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.String(http.StatusOK, name)
	})

	// form data
	e.GET("/FormData", func(c echo.Context) error {
		name := c.FormValue("name")
		return c.String(http.StatusOK, name)
	})

	e.Logger.Fatal(e.Start(":1322"))
}

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}
