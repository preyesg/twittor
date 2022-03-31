package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/preyesg/twittor.git/bd"
	"github.com/preyesg/twittor.git/jwt"
	"github.com/preyesg/twittor.git/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalido "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El Email del usuario de requerido", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o Contraseña invalido", 400)
		return
	}
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Problemas generando token de acceso "+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
