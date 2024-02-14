package models

import "github.com/faridEmilio/go-sql/database"

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(30) NOT NULL,
	email VARCHAR(30),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}

	return user
}

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.insert()
	return user
}

func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := database.Exec(sql, user.Username, user.Password, user.Email)
	user.ID, _ = result.LastInsertId()
}

//Lista todos los usuarios de la tabla
func ListUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, _ := database.Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)

		users = append(users, user)
	}

	return users
}

//Obtengo un usuario de la base de datos con su ID
func GetUser(id int) *User {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := database.Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	}
	return user
}
