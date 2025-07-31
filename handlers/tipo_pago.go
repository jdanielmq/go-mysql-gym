package handlers

import (
	"encoding/json"
	"go-mysql-gym/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetTiposPagos(rw http.ResponseWriter, r *http.Request) {
	if tiposPagos, err := models.ListTipoPago(); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, tiposPagos)
	}
}

func GetTipoPgo(rw http.ResponseWriter, r *http.Request) {
	if tipoPago, err := getTipoPagoByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		if tipoPago.IdPago == 0 {
			models.SendNoFound(rw)
		} else {
			models.SendData(rw, tipoPago)
		}
	}
	/*
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
	*/
}

func CreateTipoPago(rw http.ResponseWriter, r *http.Request) {
	tipoPago := models.TipoPago{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tipoPago); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		tipoPago.Save()
		models.SendData(rw, tipoPago)
	}
}

func UpdateTipoPago(rw http.ResponseWriter, r *http.Request) {
	var idTipoPago int64
	if tipoPago, err := getTipoPagoByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		idTipoPago = tipoPago.IdPago
	}
	tipoPago := models.TipoPago{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tipoPago); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		tipoPago.IdPago = idTipoPago
		tipoPago.Save()
		models.SendData(rw, tipoPago)
	}
}

func DeleteTipoPago(rw http.ResponseWriter, r *http.Request) {
	if tipoPago, err := getTipoPagoByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		tipoPago.Delete()
		models.SendData(rw, tipoPago)
	}
}

func getTipoPagoByRequest(r *http.Request) (models.TipoPago, error) {
	vars := mux.Vars(r)
	idTipoPago, _ := strconv.Atoi(vars["id"])
	if tipoPago, err := models.GetTipoPago(idTipoPago); err != nil {
		return *tipoPago, err
	} else {
		return *tipoPago, nil
	}
}
