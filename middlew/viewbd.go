package middlew

import (
	"net/http"

	"github.com/soUNds-unal/sounds/sounds_ms/sounds_auth_ms/bd"
)

func ViewBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ViewConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
