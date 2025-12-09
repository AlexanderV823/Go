package todo

import "fmt"

type Displayable interface {
	Display()
}

// Task структура задачи
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (t Task) Display() {
	fmt.Printf("%d: %s (выполнена: %t)", t.ID, t.Description, t.Done)
}
