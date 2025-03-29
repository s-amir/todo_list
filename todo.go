package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	name     string
	email    string
	password string
	id       int
}

type task struct {
	name       string
	duedate    time.Time
	categoryId int
	isDone     bool
	userId     int
	id         int
}
type category struct {
	id     int
	name   string
	color  string
	userId int
}

var Users []User
var Tasks []task
var isAuthenticated bool
var authenticatedUser User
var categories []category

const userFilePath string = "./users.txt"

func main2() { //callable in main
	loadUsers()
	fmt.Printf("hello to TODO app\n")
	command := flag.String("command", "noCommand", "command to rin application")
	flag.Parse()
	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin) // Create a scanner
		fmt.Print("enter another command...")
		scanner.Scan()
		*command = scanner.Text()

	}
}

func runCommand(command string) {
	if command != "register-user" && command != "exit" {
		isLoggedIn := login()
		if isLoggedIn == 0 {
			return
		}
	}

	switch command {
	case "create-task":
		createTask()
	case "register-user":
		registerUser()
	case "show-task":
		showTasks()
	case "create-category":
		createCategory()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Printf("unknown command: %s\n", command)
	}
}

func login() int {
	if isAuthenticated {
		return 1
	}

	scanner := bufio.NewScanner(os.Stdin)

	// Ask for username
	fmt.Println("Enter username for login:")
	scanner.Scan()
	name := scanner.Text()

	// Ask for password
	fmt.Println("Please enter password for login:")
	scanner.Scan()
	password := scanner.Text()

	// Check credentials
	for _, user := range Users {
		if user.name == name && user.password == password {
			isAuthenticated = true
			authenticatedUser = user
			fmt.Println("Login successful!")
			return 1
		}
	}

	// If authentication fails
	fmt.Println("Invalid username or password")
	return 0
}

func createTask() {
	var name, duedate, category string
	scanner := bufio.NewScanner(os.Stdin) // Create a scanner
	fmt.Println("Please enter a name for the task")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Please enter a name for the duedate")
	scanner.Scan()
	duedate = scanner.Text()
	parsed_time, err := time.Parse("2006-01-02", duedate)
	if err != nil {
		panic("enter valid time")
	}

	fmt.Println("Please enter a id of the category")
	scanner.Scan()
	category = scanner.Text()
	categoryId, err := strconv.Atoi(category)
	if err != nil {
		panic("faild to parse category id")
		return
	}
	var findCategory bool = false
	for _, c := range categories {
		if c.id == categoryId && c.userId == authenticatedUser.id {
			findCategory = true
		}
	}
	if !findCategory {
		panic("faild to find category")
	}

	Tasks = append(Tasks, task{name: name, duedate: parsed_time, categoryId: categoryId, isDone: false, userId: authenticatedUser.id})

	fmt.Println("done!")
}
func registerUser() {
	var id, name, email string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a name for the user")
	scanner.Scan()
	name = scanner.Text()
	fmt.Println("Please enter a name for the password")
	scanner.Scan()
	password := scanner.Text()
	fmt.Println("Please enter a email for the user")
	scanner.Scan()
	email = scanner.Text()
	id = name
	fmt.Println(id, name, email)
	u := User{name, email, password, len(Users) + 1}
	Users = append(Users, u)
	fmt.Println(Users)

	//add to file
	file, err := os.OpenFile(userFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	uString := fmt.Sprintf("Name: %s, Email: %s, Password: %s, ID: %d", u.name, u.email, u.password, u.id)
	_, err = file.WriteString(uString + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

}
func showTasks() {
	for _, task := range Tasks {
		if task.userId == authenticatedUser.id {
			fmt.Println(task.name, task.duedate, task.categoryId, task.isDone)
		}
	}
}
func createCategory() {
	var name, color string
	var id, userId int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a name for the category")
	scanner.Scan()
	name = scanner.Text()
	fmt.Println("Please enter a color for the category")
	scanner.Scan()
	color = scanner.Text()
	userId = authenticatedUser.id
	id = len(categories) + 1
	categories = append(categories, category{id, name, color, userId})

}

func loadUsers() {
	content, err := os.ReadFile(userFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var storedUsers = strings.Split(strings.TrimRight(string(content), "\n"), "\n")
	for _, user := range storedUsers {
		userFields := strings.Split(user, ",")
		var tempUser User
		for _, userField := range userFields {
			keyVaue := strings.Split(userField, ":")
			key := strings.Trim(keyVaue[0], " ")
			value := strings.Trim(keyVaue[1], " ")
			switch key {
			case "ID":
				tempUser.id, err = strconv.Atoi(value)
			case "Name":
				tempUser.name = value
			case "Email":
				tempUser.email = value
			case "Password":
				tempUser.password = value
			default:
				fmt.Println("Unknown field:", key)

			}

		}
		Users = append(Users, tempUser)
	}

}
