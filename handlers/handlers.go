package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ccmorenov/microservicesounds/middlew"
	"github.com/ccmorenov/microservicesounds/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ViewBD(routes.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
