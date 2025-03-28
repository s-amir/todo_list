package app

import "fmt"

type InMemoryStorage struct {
	UsersMap map[uint]User
}

func (storage *InMemoryStorage) CreateUser(user User) bool {
	if storage.UsersMap == nil {
		fmt.Println("UsersMap map is null")
		return false
	}
	storage.UsersMap[user.Id] = user
	return true
}
func (storage *InMemoryStorage) FindUserId(id int) User {
	return storage.UsersMap[uint(id)]
}
func (storage *InMemoryStorage) DeleteUser(id int) bool {
	if _, exists := storage.UsersMap[uint(id)]; exists && storage.UsersMap != nil {
		delete(storage.UsersMap, uint(id))
		return true
	} else {
		return false
	}

}
