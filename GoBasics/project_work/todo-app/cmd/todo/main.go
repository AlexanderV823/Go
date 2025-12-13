package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"todo-app/internal/storage"
	"todo-app/internal/todo"
)

func main() {

	var (
		cmd  string
		args []string
	)

	// Загружаем список ранее сохраненных задач из основного файла хранения
	path, err := storage.MakePath("tasks", ".json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tasks, err := storage.LoadJSON(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Считываем команду из os.Args[1] и аргументы в args := os.Args[2:]
	if len(os.Args) < 2 {
		fmt.Println("Не достаточное количество аргументов.")
		os.Exit(1)
	}
	cmd = os.Args[1]

	if len(os.Args) > 1 {
		args = os.Args[2:]
	}

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
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}

		err = storage.SaveJSON(path, tasks)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
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
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
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
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}

		err = storage.SaveJSON(path, tasks)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		deleteId := deleteCmd.Int("id", 0, "task ID")

		err = deleteCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tasks, err = todo.Delete(tasks, *deleteId)

		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}

		err = storage.SaveJSON(path, tasks)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			os.Exit(1)
		}

	case "load":
		loadCmd := flag.NewFlagSet("load", flag.ExitOnError)
		filterLoad := loadCmd.String("file", "tasks.json", "file name")

		err = loadCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		name := strings.TrimPrefix(*filterLoad, "--file=\"")
		ext := filepath.Ext(name)

		name = strings.TrimSuffix(name, ext)

		path, err = storage.MakePath(name, ext)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch ext {
		case ".json":
			tasks, err = storage.LoadJSON(path)
			if err != nil {
				fmt.Printf("Ошибка загрузки списка задач: %s\n", err)
				os.Exit(1)
			}
			path, err = storage.MakePath("tasks", ".json")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = storage.SaveJSON(path, tasks)
			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
				os.Exit(1)
			}
			fmt.Println("Список задач загружен")

		case ".csv":
			tasks, err = storage.LoadCSV(path)
			if err != nil {
				fmt.Printf("Ошибка загрузки списка задач: %s\n", err)
				os.Exit(1)
			}
			path, err = storage.MakePath("tasks", ".json")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = storage.SaveJSON(path, tasks)
			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
				os.Exit(1)
			}
			fmt.Println("Список задач загружен")
		default:
			fmt.Println("Неизвестный формат файла.")
			os.Exit(1)
		}

	case "export":

		exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
		formatExport := exportCmd.String("format", "json", "file format")
		outExport := exportCmd.String("out", "", "output file")

		err = exportCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		format := strings.TrimPrefix(*formatExport, "--format=")
		out := strings.TrimPrefix(*outExport, "--out=")

		switch format {
		case "json":

			path, err := storage.MakePath(out, "."+format)
			if err != nil {
				fmt.Printf("Ошибка сохранения списка задач: %s\n", err)
				os.Exit(1)

			}
			err = storage.SaveJSON(path, tasks)
			if err != nil {
				fmt.Printf("Ошибка сохранения списка задач: %s\n", err)
				os.Exit(1)
			}
			fmt.Println("Список задач сохранен")

		case "csv":

			path, err = storage.MakePath(out, "."+format)
			if err != nil {
				fmt.Printf("Ошибка сохранения списка задач: %s\n", err)
				os.Exit(1)
			}
			err = storage.SaveCSV(path, tasks)
			if err != nil {
				fmt.Printf("Ошибка сохранения списка задач: %s\n", err)
				os.Exit(1)
			}
			fmt.Println("Список задач сохранен")

		default:
			fmt.Println("Неизвестный формат файла.")
			os.Exit(1)
		}
	default:
		fmt.Println("Не известная команда.")
		os.Exit(1)
	}
}
