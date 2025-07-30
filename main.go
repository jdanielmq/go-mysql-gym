package main

import (
	"go-mysql-gym/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/api/tipo-pago/", handlers.GetTiposPagos).Methods("GET")
	mux.HandleFunc("/api/tipo-pago/{id:[0-9]+}", handlers.GetTipoPgo).Methods("GET")
	mux.HandleFunc("/api/tipo-pago/", handlers.CreateTipoPago).Methods("POST")
	mux.HandleFunc("/api/tipo-pago/{id:[0-9]+}", handlers.UpdateTipoPago).Methods("PUT")
	mux.HandleFunc("/api/tipo-pago/{id:[0-9]+}", handlers.DeleteTipoPago).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))
	//db.Connect()
	//fmt.Println(db.ExistsTable("inbody"))
	//medioPago := models.CreateTipoPago("Cta Rut", true)

	//db.Ping()
	//db.Close()
}
