package app

import (
	"encoding/json"
	"fmt"
	"os"
)

type App struct {
	Id              uint
	Name            string
	StorageFilePath string
}

type User struct {
	Id   uint
	Name string
}

func (app *App) CreateUser(user User) {
	var fileHandler *os.File
	if file, err := os.OpenFile(app.StorageFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666); err != nil {
		fmt.Println("can not open file", err)
	} else {
		fileHandler = file
	}
	defer func(fileHandler *os.File) {
		err := fileHandler.Close()
		if err != nil {
			fmt.Println("close file error", err)
		}
	}(fileHandler)

	mUser, mError := json.Marshal(user)
	if mError != nil {
		fmt.Println("json marshal error", mError)
	}
	_, err := fileHandler.Write(mUser)
	if err != nil {
		fmt.Println("write error", err)
		return
	}

}
