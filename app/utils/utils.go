package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"tasker-cli/app/models"
	"time"
)

var TASK_FILE_PATH string = "./tasks.json"
var CONFIG_FILE_PATH string = "./config.json"

func AddTask(taskTitle string) {
	SaveTasksToFile(taskTitle)
}

func RemoveTask(id int) {
	var tasks []models.Task
	var filteredTasks []models.Task

	_, err := GetParsedFile("{}", &tasks, TASK_FILE_PATH)

	if err != nil {
		fmt.Println("Error retrieving tasks")
	}

	for _, task := range tasks {
		if task.Id != id {
			filteredTasks = append(filteredTasks, task)
		}
	}

	StoreJsonToFile(filteredTasks)
}

func UpdateTask() {}

func GetAllTask() []models.Task {
	var tasks []models.Task

	_, err := GetParsedFile("{}", &tasks, TASK_FILE_PATH)

	if err != nil {
		fmt.Println("Error retrieving tasks")
	}

	return tasks
}

func SaveTasksToFile(title string) {
	var currentTasks []models.Task

	_, parseErr := GetParsedFile("[]", &currentTasks, TASK_FILE_PATH)

	if parseErr != nil {
		log.Fatal("Error parsing tasks from file: ", parseErr)
	}

	task := models.Task{
		Id:        getNextTaskId(),
		Title:     title,
		Status:    "pending",
		CreatedAt: time.Now().Local().String(),
		UpdatedAt: time.Now().Local().String()}

	currentTasks = append(currentTasks, task)

	StoreJsonToFile(currentTasks)

}

func StoreJsonToFile(data any) {

	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Error encoding task to JSON: ", err)
		return
	}

	filePtr, err := os.Create(TASK_FILE_PATH)

	if err != nil {
		fmt.Println("Error creating JSON file: ", err)
	}

	defer filePtr.Close()

	writeErr := os.WriteFile(TASK_FILE_PATH, []byte(jsonData), 0666)

	if writeErr != nil {
		fmt.Println("Error writing task to file:", writeErr)
		return
	}
}

func GetParsedFile[T any](fallback string, dataType *T, filePath string) (any, error) {

	// Check if file exists or not
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		err = os.WriteFile(filePath, []byte(fallback), 0644)

		if err != nil {
			return nil, fmt.Errorf("could not create file with fallback contend: %v", err)
		}
	}

	// Read file contents
	bytes, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal("Error reading tasks file: ", err)
	}

	fileText := string(bytes[:])

	//Initialise empty array string if parsed file is empty
	if len(strings.Trim(fileText, " ")) == 0 {
		fileText = fallback
	}

	parseErr := json.Unmarshal([]byte(fileText), &dataType)

	if parseErr != nil {
		return nil, parseErr
	}

	return dataType, nil
}

func getNextTaskId() int {
	var config models.Config
	var nextId int
	_, err := GetParsedFile("{}", &config, CONFIG_FILE_PATH)

	if err != nil {
		log.Fatal("Could not parse config file", err)
	}

	if config.LastInsertedId == nil {
		nextId = 1
		config.LastInsertedId = &nextId
		saveLastId(nextId)
		return nextId
	}

	nextId = *config.LastInsertedId + 1
	saveLastId(nextId)

	return nextId
}

func saveLastId(num int) {
	var config models.Config

	_, err := GetParsedFile("{}", &config, CONFIG_FILE_PATH)

	if err != nil {
		log.Fatal("Could not parse config file", err)
	}

	config.LastInsertedId = &num

	jsonData, err := json.Marshal(config)

	if err != nil {
		fmt.Println("Error encoding task to JSON: ", err)
		return
	}

	filePtr, err := os.Create(CONFIG_FILE_PATH)

	if err != nil {
		fmt.Println("Error creating JSON file: ", err)
	}

	defer filePtr.Close()

	writeErr := os.WriteFile(CONFIG_FILE_PATH, []byte(jsonData), 0666)

	if writeErr != nil {
		fmt.Println("Error writing task to file:", writeErr)
		return
	}
}

// func GetParsedDate (){
// 	timeStr := "2025-05-03 15:06:05.586395655 +0530 IST"

// 	// Use time.RFC3339 or the exact layout from time.String()
// 	layout := "2006-01-02 15:04:05.999999999 -0700 MST"

// 	parsedTime, err := time.Parse(layout, timeStr)
// 	if err != nil {
// 		fmt.Println("Error parsing time:", err)
// 	} else {
// 		fmt.Println("Parsed time:", parsedTime)
// 	}
// }
