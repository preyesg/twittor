package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/preyesg/twittor.git/models"
)

func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("MiclavePreyesg")

	paylaop := jwt.MapClain{
		"email":           t.Email,
		"nombre":          t.Nombre,
		"apellidos":       t.Apellidos,
		"fechaNacimiento": t.FechaNacimiento,
		"id":              t.ID.Hex(),
		"exp":             time.Hour.Add(time.Hour * 24).UniX(),
	}

	token := jwt.NewWithClaims(jwt.SigninMethodHS256, paylaop)

	tokenStr := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
