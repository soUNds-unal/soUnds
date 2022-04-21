package routes

import (
	"encoding/json"
	"net/http"

	"github.com/soUNds-unal/sounds/sounds_ms/sounds_auth_ms/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Falta el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.SearchPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro "+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
