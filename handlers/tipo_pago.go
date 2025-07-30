package handlers

import (
	"encoding/json"
	"fmt"
	"go-mysql-gym/db"
	"go-mysql-gym/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetTiposPagos(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	db.Connect()
	tiposPagos := models.ListTipoPago()
	db.Close()
	output, _ := json.Marshal(tiposPagos)
	fmt.Fprintln(rw, string(output))
}

func GetTipoPgo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	db.Connect()
	tipoPago := models.GetTipoPago(id)
	db.Close()
	if tipoPago.IdPago == 0 {
		rw.WriteHeader(http.StatusNoContent)
		fmt.Fprintln(rw, "")
	} else {
		output, _ := json.Marshal(tipoPago)
		fmt.Fprintln(rw, string(output))
	}
}

func CreateTipoPago(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "creando un nuevo tipo de pago")
}

func UpdateTipoPago(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "actualizando un tipo de pago segun el id")
}

func DeleteTipoPago(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "elimimando un tipo de pago")
}
