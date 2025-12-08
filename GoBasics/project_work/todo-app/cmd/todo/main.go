package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"todo-app/internal/storage"
	"todo-app/internal/todo"
)

func clearConsole() {
	var cmd *exec.Cmd

	// Определяем операционную систему
	if runtime.GOOS == "windows" {
		// Для Windows используется команда 'cls'
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		// Для Linux и macOS используется команда 'clear'
		cmd = exec.Command("clear")
	}
	// Устанавливаем стандартный вывод (терминал) как вывод команды
	cmd.Stdout = os.Stdout
	// Выполняем команду
	cmd.Run()
}

func main() {
	fmt.Println("Мини-приложение для управления задачами на Go")

	// Загружаем список ранее сохраненных задач из основного файла хранения
	path, err := storage.MakePath("tasks.json")
	if err != nil {
		fmt.Printf("Ошибка: %s", err)
		os.Exit(1)
	}
	tasks, err := storage.LoadJSON(path)

	if err != nil {
		fmt.Printf("Ошибка: %s", err)
		os.Exit(1)
	}

	for {
		fmt.Print("Введите команду: ")

		// Очищаем консоль перед вводом
		clearConsole()

		// Считываем команду из os.Args[1] и аргументы в args := os.Args[2:]
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

			desc := strings.TrimPrefix(*descAdd, "--desc=\"")
			desc = strings.TrimSuffix(desc, "\"")

			tasks, err = todo.Add(tasks, desc)

			if err != nil {
				fmt.Printf("Ошибка: %s", err)
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

			filter := strings.TrimPrefix(*filterList, "--filter=")

			filteredTasks, err := todo.List(tasks, filter)

			if err != nil {
				fmt.Printf("Ошибка: %s", err)
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
				fmt.Printf("Ошибка: %s", err)
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
				fmt.Printf("Ошибка: %s", err)
				continue
			}

		case "load":
			loadCmd := flag.NewFlagSet("load", flag.ExitOnError)
			filterLoad := loadCmd.String("file", "tasks.json", "file name")

			path := strings.TrimPrefix(*filterLoad, "--file=\"")

			ext := filepath.Ext(path)

			switch ext {
			case "json":
				tasks, err = storage.LoadJSON(path)
				if err != nil {
					fmt.Printf("Ошибка загрузки списка задач: %s", err)
					continue
				}
				fmt.Println("Список задач загружен")
			case "csv":
				tasks, err = storage.LoadCSV(path)
				if err != nil {
					fmt.Printf("Ошибка загрузки списка задач: %s", err)
					continue
				}
				fmt.Println("Список задач загружен")
			default:
				fmt.Println("Неизвестный формат файла.")
			}

		case "export":

			exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
			formatExport := exportCmd.String("format", "json", "file format")
			outExport := exportCmd.String("out", "", "output file")

			format := strings.TrimPrefix(*formatExport, "--format=")
			out := strings.TrimPrefix(*outExport, "--out=")

			switch format {
			case "json":

				path, err := storage.MakePath(out + "." + format)
				if err != nil {
					fmt.Printf("ООшибка сохранения списка задач: %s", err)
					continue
				}
				err = storage.SaveJSON(path, tasks)
				if err != nil {
					fmt.Printf("Ошибка сохранения списка задач: %s", err)
					continue
				}
				fmt.Println("Список задач сохранен")

			case "csv":

				path, err := storage.MakePath(out + "." + format)
				if err != nil {
					fmt.Printf("Ошибка сохранения списка задач: %s", err)
					continue
				}
				err = storage.SaveCSV(path, tasks)
				if err != nil {
					fmt.Printf("Ошибка сохранения списка задач: %s", err)
					continue
				}
				fmt.Println("Список задач сохранен")

			default:
				fmt.Println("Неизвестный формат файла.")
			}

		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}
