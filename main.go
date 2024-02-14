package main

import (
	database "github.com/faridEmilio/go-sql/database"
	"github.com/faridEmilio/go-sql/pkg/models"
)

func main() {
	database.ConnectDB()
	database.CreateTable(models.UserSchema, "users")
	database.CLoseDB()
}
