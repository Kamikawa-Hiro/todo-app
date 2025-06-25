package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const dataFile = "todo.json"

func loadTasks() ([]Task, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(tasks)
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func AddTask(title string) {
	tasks, _ := loadTasks()

	newTask := Task{
        ID:    nextID(tasks),
        Title: title,
        Done:  false,
    }
 
    tasks = append(tasks, newTask)
 
    if err := saveTasks(tasks); err != nil {
        fmt.Println("Error saving tasks:", err)
        return
    }
 
    fmt.Printf("Added task: \"%s\"\n", title)
}

func ListTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error")
	}

	for _, t := range tasks {
		if t.Done == true {
			fmt.Printf("%d: %s [x]\n", t.ID, t.Title)
		}else {
			fmt.Printf("%d: %s [ ]\n", t.ID, t.Title)
		}	
	}
}

func CompleteTask(id int) {
	panic("unimplemented")
}

func DeleteTask(id int) {
	panic("unimplemented")
}