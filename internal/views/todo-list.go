package views

import (
	"htmx-test/internal/todo"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ToDoListView struct {
	ToDos todo.ToDoList
}

func ToDoList(tdl todo.ToDoList) func(echo.Context) error {
	view := func(c echo.Context) error {
		return c.Render(http.StatusOK, "todo-list.html", ToDoListView{
			ToDos: tdl,
		})
	}

	return view
}
