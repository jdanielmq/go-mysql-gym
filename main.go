package main

import (
	"fmt"
	"go-mysql-gym/db"
	"go-mysql-gym/models"
	"time"
)

func main() {
	db.Connect()
	fmt.Println(db.ExistsTable("inbody"))
	//medioPago := models.CreateTipoPago("Cta Rut", true)

	inbodys := models.ListInbodys()
	fmt.Println(inbodys)

	timecurrent := (time.Now()).Format(time.RFC3339Nano)
	fmt.Println(timecurrent)

	//inbody := models.CreateInbody("14125469-3", false, "{\"key\": \"value\"}", timecurrent, "comentarios", 2)
	//fmt.Println(inbody)

	//inbodys = models.ListInbodys()
	//fmt.Println(inbodys)
	//db.Ping()
	db.Close()
}
