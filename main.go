package main

import (
	"log"

	"twittor/bd"
	"twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal(("Base sin conexion"))
		return
	}
	handlers.Manejadores()

}
