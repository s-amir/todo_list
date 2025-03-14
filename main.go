package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	name  string
	email string
	id    string
}

var Users []User

func main() {
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

	switch command {
	case "create-task":
		createTask()
	case "register-user":
		registerUser()
	case "exit":
		os.Exit(0)
	default:
		fmt.Printf("unknown command: %s\n", command)
	}
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

	fmt.Println("Please enter a name for the category")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println(name, duedate, category)
}
func registerUser() {
	var id, name, email string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a name for the user")
	scanner.Scan()
	name = scanner.Text()
	fmt.Println("Please enter a email for the user")
	scanner.Scan()
	email = scanner.Text()
	id = name
	fmt.Println(id, name, email)
	Users = append(Users, User{name, email, id})
	fmt.Println(Users)

}
