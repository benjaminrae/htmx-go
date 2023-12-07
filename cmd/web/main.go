package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"htmx-test/internal/todo"
	"htmx-test/internal/views"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	toDoList := todo.ToDoList{
		ToDos: []todo.ToDo{
			{
				Task:       "Do something",
				Id:         1,
				IsComplete: false,
			},
		},
	}

	e.GET("/todos", views.ToDoList(toDoList))
	e.POST("/todos", func(c echo.Context) error {
		toDoList.CreateTodo(c.FormValue("task"))
		fmt.Println(toDoList)
		return c.Render(http.StatusOK, "todo-list.html", views.ToDoListView{
			ToDos: toDoList,
		})
	})

	e.Static("/css/", "web/css")
	e.Static("/assets/", "web/assets")
	e.Static("/js/", "web/js")
	e.GET("/ping", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
