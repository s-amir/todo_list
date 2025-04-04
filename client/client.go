package main

import (
	"awesomeProject/delievery/delieveryparam"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("command", os.Args[0])
	message := "default message"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	conn, err := net.Dial("tcp", "127.0.0.1:1993")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	req := delieveryparam.Request{Command: message}
	if req.Command == "create-task" {
		req.CreateTaskRequest = delieveryparam.CreateTaskRequest{
			Title:      "Task1",
			DueDate:    "test",
			CategoryID: 1,
		}
	}
	serializedData, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	numberOfBytes, wErr := conn.Write(serializedData)
	if wErr != nil {
		log.Fatal(wErr)
	}
	fmt.Printf("wrote %d bytes\n", numberOfBytes)
	var data = make([]byte, 1024)
	numberOfRead, readErr := conn.Read(data)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Printf("response is : %s\n", string(data[:numberOfRead]))
}
