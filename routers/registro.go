package routers

import (
	"encoding/json"
	"net/http"

	"github.com/preyesg/twittor.git/bd"
	"github.com/preyesg/twittor.git/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos enviados "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 5 {
		http.Error(w, "Contraseña de 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Usuario ya registrado", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Error insertando registro "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Problemas insertando registro", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
