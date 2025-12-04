package todo

import "fmt"

type Displayable interface {
	Display()
}

// Task структура задачи
type Task struct {
	ID          int
	Description string
	Done        bool
}

func (t Task) Display() {
	fmt.Printf("%d: %s (выполнена: %t)", t.ID, t.Description, t.Done)
}
