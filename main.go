package main

import (
	"fmt"

	database "github.com/faridEmilio/go-sql/database"
	"github.com/faridEmilio/go-sql/pkg/models"
)

func main() {
	database.ConnectDB()
	//database.CreateTable(models.UserSchema, "users")
	//models.CreateUser("emilio", "2451", "emilio@gmail.com")

	//users := models.ListUsers()
	//fmt.Println(users)

	user := models.GetUser(1)
	fmt.Println(user)
	database.CLoseDB()
}
