package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"todo-app/internal/storage"
	"todo-app/internal/todo"
)

func main() {
	fmt.Println("Мини-приложение для управления задачами на Go")

	//Подгрузите текущие задачи
	//Файл хранения задаем сами
	//Загружаем список ранее сохраненных задач из основного файла хранения
	dataFile := "tasks.json"
	tasks, err := storage.LoadJSON(dataFile)

	for {

		//Считываем команду из os.Args[1] и аргументы в args := os.Args[2:]
		fmt.Print("Введите команду: ")
		cmd := os.Args[1]
		args := os.Args[2:]

		switch cmd {
		case "add":
			addCmd := flag.NewFlagSet("add", flag.ExitOnError)
			descAdd := addCmd.String("desc", "", "task description")

			err = addCmd.Parse(args)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			tasks, err = todo.Add(tasks, *descAdd)

			if err != nil {
				fmt.Println("Ошибка: %s", err)
				continue
			}

		case "list":
			listCmd := flag.NewFlagSet("list", flag.ExitOnError)
			filterList := listCmd.String("filter", "all", "all, done, pending")

			err = listCmd.Parse(args)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			filteredTasks, err := todo.List(tasks, *filterList)

			if err != nil {
				fmt.Println("Ошибка: %s", err)
				continue
			}

			for _, t := range filteredTasks {
				t.Display()
			}

		case "complete":
			completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
			completeId := completeCmd.Int("id", 0, "task ID")

			err = completeCmd.Parse(args)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			tasks, err = todo.Complete(tasks, *completeId)

			if err != nil {
				fmt.Println("Ошибка: %s", err)
				continue
			}

		case "delete":
			deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
			deleteId := deleteCmd.Int("id", 0, "task ID")

			err = deleteCmd.Parse(args)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			tasks, err = todo.Complete(tasks, *deleteId)

			if err != nil {
				fmt.Println("Ошибка: %s", err)
				continue
			}

		case "load":
			loadCmd := flag.NewFlagSet("load", flag.ExitOnError)
			filterLoad := loadCmd.String("file", "tasks.json", "file name")

			// ext := filepath.Ext()
			// err := storage.LoadJSON(cmd)
			// if err != nil {
			// 	fmt.Printf("Ошибка загрузки списка задач: %s\n", err)
			// 	continue
			// }
			// fmt.Println("Список задач загружен")

		case "export":
			exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
			formatExport := exportCmd.String("format", "json", "file format")
			outExport := exportCmd.String("out", "tasks.json", "output format")

		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}
