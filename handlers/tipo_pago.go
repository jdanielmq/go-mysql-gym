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
	tiposPagos, _ := models.ListTipoPago()
	db.Close()
	output, _ := json.Marshal(tiposPagos)
	fmt.Fprintln(rw, string(output))
}

func GetTipoPgo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	db.Connect()
	tipoPago, _ := models.GetTipoPago(id)
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
	rw.Header().Set("Content-Type", "application/json")
	tipoPago := models.TipoPago{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tipoPago); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		tipoPago.Save()
		db.Close()
	}
	output, _ := json.Marshal(tipoPago)
	fmt.Fprintln(rw, string(output))
}

func UpdateTipoPago(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	tipoPago := models.TipoPago{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tipoPago); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		tipoPago.Save()
		db.Close()
	}
	output, _ := json.Marshal(tipoPago)
	fmt.Fprintln(rw, string(output))
}

func DeleteTipoPago(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idTipoPago, _ := strconv.Atoi(vars["id"])
	db.Connect()
	tipoPago, _ := models.GetTipoPago(idTipoPago)
	tipoPago.Delete()
	db.Close()
	output, _ := json.Marshal(tipoPago)
	fmt.Fprintln(rw, string(output))
}
