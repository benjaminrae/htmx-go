package todo

type ToDo struct {
	Id         int
	Task       string
	IsComplete bool
}

type ToDoList struct {
	ToDos []ToDo
}

func (tdl *ToDoList) CreateTodo(task string) {
	var nextId = tdl.getNextId()

	if len(tdl.ToDos) == 0 {
		tdl.ToDos = []ToDo{}
	}

	var toDo = ToDo{
		Id:         nextId,
		IsComplete: false,
		Task:       task,
	}

	tdl.ToDos = append(tdl.ToDos, toDo)
}

func (tdl *ToDoList) DeleteTodo(taskId int) {
	for i, toDo := range tdl.ToDos {
		if toDo.Id == taskId {
			tdl.ToDos = append(tdl.ToDos[:i], tdl.ToDos[i+1:]...)
		}
	}
}

func (tdl *ToDoList) getNextId() int {
	if len(tdl.ToDos) == 0 {
		return 1
	}

	var highest = 1

	for _, toDo := range tdl.ToDos {
		if toDo.Id > highest {
			highest = toDo.Id
		}
	}

	return highest + 1
}
