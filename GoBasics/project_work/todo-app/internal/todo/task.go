package todo

import "fmt"

type Displayable interface {
	Display()
}

// Task структура задачи
type Task struct {
	ID int
	Description string
	Done bool
}

func (t Task) Display() {
	fmt.Printf("ID: %d\nЗадача: %s\nСтатус: %t", t.ID, t.Description, t.Done)
}