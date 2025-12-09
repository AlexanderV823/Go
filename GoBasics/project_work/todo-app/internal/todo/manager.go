package todo

import "fmt"

// Add добавляет задачу в список.
// tasks - срез структур с задачами,
// desc  - описание задачи
func Add(tasks []Task, desc string) ([]Task, error) {

	var taskId int = 1

	// Ищем свободный id в диапазоне от 1 до 250
	for id := 1; id < 251; id++ {
		// Проверяем наличие в срезе с задачами ранее добавленной задачи с таким же id
		for _, v := range tasks {
			if v.ID == id {
				// Если такой id ранее было присвоено какой-то из загруженных задач, то переходим к новой генерации
				continue
			} else {
				// При нахождении первого ранее не использованного id прерываем цикл
				taskId = id
				break
			}
		}
	}

	newTask := Task{ID: taskId, Description: desc, Done: false}

	tasks = append(tasks, newTask)

	fmt.Printf("Задача №%d добавлена в список\n", taskId)

	return tasks, nil
}

// List выводит список задач.
// tasks  - срез структур с задачами,
// filter - all (по умолчанию) — все задачи;
//   - done — только выполненные;
//   - pending — только невыполненные.
func List(tasks []Task, filter string) ([]Task, error) {

	var filteredTasks []Task

	switch filter {
	case "all":
		return tasks, nil
	case "done":
		for _, t := range tasks {
			if t.Done {
				filteredTasks = append(filteredTasks, t)
			}
		}
		return filteredTasks, nil
	case "pending":
		for _, t := range tasks {
			if !t.Done {
				filteredTasks = append(filteredTasks, t)
			}
		}
		return filteredTasks, nil
	default:
		return tasks, fmt.Errorf("передан не действительный фильтр")
	}
}

// Complete устанавливает отметку выполнения.
// tasks - срез структур с задачами,
// id    - ID задачи
func Complete(tasks []Task, id int) ([]Task, error) {

	var flag bool

	for i, v := range tasks {
		if v.ID == id {
			tasks[i].Done = true
			flag = true
		}
	}

	if flag {
		fmt.Printf("Задача №%d отмечена как выполненная\n", id)
	} else {
		fmt.Printf("Задача №%d не найдена в списке\n", id)
	}

	return tasks, nil
}

// Delete удаляет задачу.
// tasks - срез структур с задачами,
// id    - ID задачи
func Delete(tasks []Task, id int) ([]Task, error) {

	var flag bool

	for i, v := range tasks {
		if v.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			flag = true
		}
	}

	if flag {
		fmt.Printf("Задача №%d отмечена как выполненная\n", id)
	} else {
		fmt.Printf("Задача №%d не найдена в списке\n", id)
	}

	return tasks, nil
}
