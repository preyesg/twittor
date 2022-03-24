package handlers

import (
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func MAnejadores() {

	router := mux.NewRouter()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handlers := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handlers))

}
