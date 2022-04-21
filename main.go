package main

import (
	"log"

	"github.com/soUNds-unal/sounds/sounds_ms/sounds_auth_ms/bd"
	"github.com/soUNds-unal/sounds/sounds_ms/sounds_auth_ms/handlers"
)

func main() {
	if bd.ViewConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
