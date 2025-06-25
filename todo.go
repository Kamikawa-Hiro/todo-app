package main

import (
	"encoding/json"
	"os"
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
	panic("unimplemented")
}

func ListTasks() {
	panic("unimplemented")
}

func CompleteTask(id int) {
	found := false
	tasks, err := loadTasks()

	for i := range tasks {
		if tasks[i].ID == targetID {
			tasks[i].Done = true
			found = true
			break // 見つけたらループ終了
		}
	}
	if !found {
		return errors.New("task not found")
	}

	err := saveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	} else {
		fmt.Println("Tasks saved successfully!")
	}
}

func DeleteTask(id int) {
	found := false
	newTasks := make([]Task, 0, len(tasks)) // 新しいスライスを作る
	tasks, err := loadTasks()
	
	for _, task := range tasks {
		if task.ID == targetID {
			found = true
			continue // この task はスキップ（削除）
		}
		newTasks = append(newTasks, task)
	}

	if !found {
		return tasks, errors.New("task not found")
	}

	err := saveTasks(newTasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	} else {
		fmt.Println("Tasks saved successfully!")
	}
}
