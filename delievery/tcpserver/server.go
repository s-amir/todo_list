package main

import (
	"awesomeProject/delievery/delieveryparam"
	"awesomeProject/server/repository/inmem"
	"awesomeProject/server/service/task"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

const network = "tcp"
const addr = "127.0.0.1:1993"

func main() {
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Println("listening on", listener.Addr())

	memTask := inmem.NewTask()
	taskService := task.NewTaskService(memTask)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		rawRequest := make([]byte, 1024)
		numberOfBytes, err := conn.Read(rawRequest)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("client address is : %s , number of read bytes is : %d , the message is : %s \n", conn.RemoteAddr(), numberOfBytes, string(rawRequest))

		req := delieveryparam.Request{}
		err = json.Unmarshal(rawRequest[:numberOfBytes], &req)
		if err != nil {
			log.Println(err)
			continue
		}

		switch req.Command {
		case "create-task":
			response, err := taskService.Create(task.CreateRequest{
				Title:               req.CreateTaskRequest.Title,
				DueDate:             req.CreateTaskRequest.DueDate,
				CategoryID:          req.CreateTaskRequest.CategoryID,
				AuthenticatedUserID: 0,
			})
			if err != nil {
				log.Println(err)
				continue
			}
			data, err := json.Marshal(response)
			if err != nil {
				log.Println(err)
				continue
			}
			_, err = conn.Write(data)
			if err != nil {
				log.Println(err)
				continue
			}

		}
		conn.Close()

	}
}
