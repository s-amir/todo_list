package app

import (
	"errors"
	"fmt"
	"time"
)
import "awesomeProject/logs"

type InMemoryStorage struct {
	UsersMap map[uint]User
}

func (storage *InMemoryStorage) CreateUser(user User) (bool, error) {
	if storage.UsersMap == nil {
		fmt.Println("UsersMap map is null")
		return false, &logs.RichError{Message: "UsersMap is null", MetaData: nil, Operation: "CreateUser", Time: time.Now()}
	}
	storage.UsersMap[user.Id] = user
	return true, nil
}

func (storage *InMemoryStorage) FindUserId(id int) User {
	return storage.UsersMap[uint(id)]
}

func (storage *InMemoryStorage) DeleteUser(id int) (bool, error) {
	if _, exists := storage.UsersMap[uint(id)]; exists && storage.UsersMap != nil {
		delete(storage.UsersMap, uint(id))
		return true, nil
	} else {
		//return false, &logs.RichError{Message: "User not found", MetaData: nil, Operation: "DeleteUser", Time: time.Now()}
		return false, errors.New("user not found")
	}

}
