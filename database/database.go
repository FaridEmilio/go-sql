package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "root:admin@tcp(localhost:3306)/go"

var db *sql.DB

func ConnectDB() {
	connection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Conexión exitosa")
	db = connection
}

func CLoseDB() {
	db.Close()
}

// Verifica la conexión
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// CREA UNA NUEVA TABLA
func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// VERIFICAR EXISTENCIA DE TABLA
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

// Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return result, err
}

// Polimorfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return rows, err
}
