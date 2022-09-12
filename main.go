package main

import (
	"log"

	"github.com/franciscolivam/twittor/bd"
	"github.com/franciscolivam/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
