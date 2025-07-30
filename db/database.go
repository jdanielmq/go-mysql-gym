package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "jdanielmq:JdmQ1481@tcp(localhost:3306)/db_grossgym_fitness"

var db *sql.DB

func Connect() {
	connection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion Exitosa.!!!")
	db = connection
}

func Close() {
	db.Close()
}

// verifica que la base de datos exista
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return rows.Next()
}

// funcion que crea un tabla, segun el string que se le pasa
func CreateTable(shema string, nameTable string) {
	if !ExistsTable(nameTable) {
		_, err := db.Exec(shema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// reiniciar el registro de una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

// polimorfismo de exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

// polimorfismo de query
func Query(query string, args ...any) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
