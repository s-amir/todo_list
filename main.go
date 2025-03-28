package main

import (
	"awesomeProject/app"
	"fmt"
)

func main() {
	//test.DayOfWeek(1)
	//-----------------------------------------
	user1 := app.User{
		Id:   uint(1),
		Name: "user1",
	}
	user2 := app.User{
		Id:   uint(2),
		Name: "user2",
	}

	user3 := app.User{
		Id:   uint(3),
		Name: "user3",
	}

	users := map[uint]app.User{
		user1.Id: user1,
		user2.Id: user2,
		user3.Id: user3,
	}
	app1 := app.App{
		Id:   1,
		Name: "app1",
		UserStorage: &app.InMemoryStorage{
			UsersMap: users,
		},
	}
	fmt.Println(app1.FindUserId(1))
	fmt.Println(app1.CreateUser(app.User{Id: 4, Name: "user4"}))
	fmt.Println(app1.FindUserId(4))
	fmt.Println(app1.DeleteUser(1))
	fmt.Println(app1.FindUserId(2))
}
