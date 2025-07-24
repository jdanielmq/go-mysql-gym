package main

import (
	"fmt"
	"go-mysql-gym/db"
	"go-mysql-gym/models"
)

func main() {
	db.Connect()
	fmt.Println(db.ExistsTable("estado"))
	//medioPago := models.CreateTipoPago("Cta Rut", true)
	estados := models.ListEstados()
	fmt.Println(estados)

	estado := models.CreateEstado("Habilitado", true)
	fmt.Println(estado)

	estado = models.CreateEstado("DesaHabilitado", true)
	fmt.Println(estado)

	estado = models.CreateEstado("Pendiente", true)
	fmt.Println(estado)

	estado = models.CreateEstado("Suspendido", false)
	fmt.Println(estado)

	estados = models.ListEstados()
	fmt.Println(estados)

	//db.Ping()
	db.Close()
}
