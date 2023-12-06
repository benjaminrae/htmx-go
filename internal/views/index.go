package views

import (
	"htmx-test/internal/todo"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexPage struct {
	ToDos todo.ToDoList
}

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", IndexPage{
		ToDos: todo.ToDoList{
			ToDos: []todo.ToDo{},
		},
	})
}
