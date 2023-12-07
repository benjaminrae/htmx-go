package views

import (
	"htmx-test/internal/todo"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexPage struct {
	ToDoList todo.ToDoList
}

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", IndexPage{
		ToDoList: todo.ToDoList{
			ToDos: []todo.ToDo{
				{Task: "Do something", Id: 1, IsComplete: false},
			},
		},
	})
}
