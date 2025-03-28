package app

// StorageLayer use dependency inversion with logs
type StorageLayer interface {
	CreateUser(user User) (bool, error)
	FindUserId(id int) User
	DeleteUser(id int) (bool, error)
}

type App struct {
	Id   uint
	Name string
	//StorageFilePath string
	UserStorage StorageLayer
}

type User struct {
	Id   uint
	Name string
}

//func (app *App) CreateUser(user User) {
//	var fileHandler *os.File
//	if file, err := os.OpenFile(app.StorageFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666); err != nil {
//		fmt.Println("can not open file", err)
//	} else {
//		fileHandler = file
//	}
//	defer func(fileHandler *os.File) {
//		err := fileHandler.Close()
//		if err != nil {
//			fmt.Println("close file error", err)
//		}
//	}(fileHandler)
//
//	mUser, mError := json.Marshal(user)
//	if mError != nil {
//		fmt.Println("json marshal error", mError)
//	}
//	_, err := fileHandler.Write(mUser)
//	if err != nil {
//		fmt.Println("write error", err)
//		return
//	}

//}

func (app *App) CreateUser(user User) (bool, error) {
	return app.UserStorage.CreateUser(user)
}
func (app *App) FindUserId(id int) User {
	return app.UserStorage.FindUserId(id)
}
func (app *App) DeleteUser(id int) (bool, error) {
	return app.UserStorage.DeleteUser(id)
}
