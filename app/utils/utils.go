package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"tasker-cli/app/models"
)

var TASK_FILE_PATH string = "./tasks.json"

func AddTask(taskTitle string) {
	SaveTasksToFile(taskTitle)
}

func RemoveTask() {}

func UpdateTask() {}

func GetAllTask() {}

func SaveTasksToFile(title string) {
	task := models.Task{Id: 1, Title: title, Status: "pending", CreatedAt: "", UpdatedAt: ""}

	jsonData, err := json.Marshal(task)

	if err != nil {
		fmt.Println("Error encoding task to JSON:", err)
		return
	}

	filePtr, err := os.Create(TASK_FILE_PATH)

	if err != nil {
		fmt.Println("Error creating JSON file:", err)
	}

	defer filePtr.Close()

	// writeErr := os.WriteFile(TASK_FILE_PATH, []byte(jsonData), 0666)

	// if writeErr != nil {
	// 	fmt.Println("Error writing task to file:", writeErr)
	// 	return
	// }

	fmt.Println(jsonData)
}

func GetParsedTaskList() {}
