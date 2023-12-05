package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"htmx-test/internal/views"
)

// https://echo.labstack.com/docs/templates#rendering
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("web/**/*.html")),
	}
	e.Renderer = t

	e.Use(middleware.Logger())

	e.GET("/", views.Index)

	e.GET("/ping", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
