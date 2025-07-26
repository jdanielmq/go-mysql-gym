package main

import (
	"fmt"
	"go-mysql-gym/db"
	"go-mysql-gym/models"
)

func main() {
	db.Connect()
	fmt.Println(db.ExistsTable("instructor"))
	//medioPago := models.CreateTipoPago("Cta Rut", true)
	instructores := models.ListInstructores()
	fmt.Println(instructores)

	instructor := models.CreateInstructor("Ayax GrossPellier", true)
	fmt.Println(instructor)

	instructor = models.CreateInstructor("Gaston GrossPellier", true)
	fmt.Println(instructor)

	instructor = models.CreateInstructor("Luis Quezadas", true)
	fmt.Println(instructor)

	instructores = models.ListInstructores()
	fmt.Println(instructores)
	//db.Ping()
	db.Close()
}
