package routes

import (
	"net/http"

	"github.com/soUNds-unal/sounds/sounds_ms/sounds_auth_ms/bd"
)

func EliminarUser(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar parametro de ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroUser(IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar al usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
