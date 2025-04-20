package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID           int
	Title        string
	Description  string
	RegisterDate time.Time
	FinishedDate time.Time
	Completed    bool
}

type TaskManager struct {
	Tasks     []*Task
	TasksByID map[int]*Task
}

// --- String Template For Tasks

var tmpl, _ = template.New("printTask").Parse("----------\nTask ID: {{.ID}}\nTitle: {{.Title}}\nDescription: {{.Description}}\nRegistered at: {{.RegisterDate}}\nFinished at: {{.FinishedDate}}\nCompleted: {{.Completed}}\n")

// --- TaskManager Methods

func (tm *TaskManager) AddTask(t *Task) {
	tm.Tasks = append(tm.Tasks, t)
	tm.TasksByID[t.ID] = t
	fmt.Println("Task added successfully!")
}

func (tm *TaskManager) CompleteTask(t *Task) {
	t.Completed = true
	t.FinishedDate = time.Now()
	fmt.Println("Finished at", t.FinishedDate.Format("02-01-06 15-04"))
}

func (tm *TaskManager) ListTasks() {
	for _, v := range tm.TasksByID {
		tmpl.Execute(os.Stdout, v)
	}
}

func (tm *TaskManager) FindByID(id int) (*Task, error) {
	for k, v := range tm.TasksByID {
		if k == id {
			return v, nil
		}
	}
	return &Task{}, TaskNotFoundError{"Task not found."}
}

func (tm *TaskManager) AddMultipleTasks(t ...Task) {
	for _, v := range t {
		tm.Tasks = append(tm.Tasks, &v)
		tm.TasksByID[v.ID] = &v
	}
	fmt.Println("Tasks added successfully!")
}

// --- Custom Tasks Error
type TaskNotFoundError struct {
	message string
}

func (err TaskNotFoundError) Error() string {
	return err.message
}

// --- ID Gen
func ClosureGen() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

var GenerateID = ClosureGen()

// Initializing a Reader
var reader = bufio.NewReader(os.Stdin)

func main() {
	defer fmt.Println("Thanks for using To Do Go!")
	// Initializing a TaskManager
	manager := TaskManager{
		Tasks:     []*Task{},
		TasksByID: map[int]*Task{},
	}

	for {
		PrintMenu()
		fmt.Print("Select a choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			newTask := ReturnTask()
			manager.AddTask(&newTask)
		case "2":
			multiTask := ReturnMultipleTasks()
			manager.AddMultipleTasks(multiTask...)
		case "3":

			delete(manager.TasksByID, 0)
			manager.ListTasks()
		case "4":
			fmt.Print("Write the task ID: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)

			id, err := strconv.Atoi(idStr)

			if err != nil {
				fmt.Println("Invalid ID...")
				continue
			}

			findedTask, err := manager.FindByID(id)

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			tmpl.Execute(os.Stdout, findedTask)
		case "5":
			fmt.Print("Write the task ID: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)

			id, err := strconv.Atoi(idStr)

			if err != nil {
				fmt.Println("Invalid ID...")
				continue
			}
			findedTask, err := manager.FindByID(id)

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			manager.CompleteTask(findedTask)
		case "6":
			return
		}
	}
}

// --- Printing Menu
func PrintMenu() {
	fmt.Println("\n******************")
	fmt.Println("*TO DO GO PROGRAM*")
	fmt.Println("******************")

	fmt.Println("1. Add New Task")
	fmt.Println("2. Add Multiple Tasks")
	fmt.Println("3. List All Tasks")
	fmt.Println("4. Find Task By ID")
	fmt.Println("5. Complete Task")
	fmt.Println("6. Exit\n")
}

func RequestTaskData() (int, string, string, time.Time, bool) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	var title, description string
	fmt.Print("Write the title of the task: ")
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Print("Write the description of the task: ")
	description, _ = reader.ReadString('\n')
	description = strings.TrimSpace(description)

	if title == "" {
		panic("title cant be empty!")
	}

	return GenerateID(), title, description, time.Now(), false
}

func ReturnTask() Task {
	id, title, description, registeredDate, completed := RequestTaskData()
	return Task{
		ID:           id,
		Title:        title,
		Description:  description,
		RegisterDate: registeredDate,
		Completed:    completed,
	}
}

func ReturnMultipleTasks() []Task {
	var total int
	var multiTasks []Task
	fmt.Print("How many tasks do you want to register? ")
	fmt.Scanln(&total)

	for i := 0; i < total; i++ {
		fmt.Printf("Task nº %d\n", i+1)
		id, title, description, registeredDate, completed := RequestTaskData()
		multiTasks = append(multiTasks, Task{ID: id, Title: title, Description: description, RegisterDate: registeredDate, Completed: completed})
	}
	return multiTasks
}
