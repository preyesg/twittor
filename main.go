package main

import (
	"log"

	"github.com/preyesg/twittor.git/bd"
	"github.com/preyesg/twittor.git/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal(("Base sin conexion"))
		return
	}
	handlers.Manejadores()

}
