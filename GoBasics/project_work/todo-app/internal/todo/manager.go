package todo

import (
	"fmt"
)

// Add добавляет задачу в список
// tasks - срез структур с задачами
// desc - описание задачи
func Add (tasks []Task, desc string) ([]Task, error) {
fmt.Println("Add добавляет задачу в список")
return nil, nil
}

// List выводит список задач
// tasks - срез структур с задачами
// filter - all (по умолчанию) — все задачи;
//        - done — только выполненные;
//        - pending — только невыполненные.
func  List(tasks []Task, filter string) ([]Task, error) {
fmt.Println("List выводит список задач")
return nil, nil
}

// Complete выводит список задач
// tasks - срез структур с задачами
// id - ID задачи
func Complete(tasks []Task, id int) ([]Task, error) {
fmt.Println("Complete выводит список задач")
return nil, nil
}

// Delete удаляет задачу
// tasks - срез структур с задачами
// id - ID задачи
func Delete(tasks []Task, id int) ([]Task, error) {
fmt.Println("Delete удаляет задачу")
return nil, nil
}