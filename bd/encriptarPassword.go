package bd

import (
	"github.com/golang/crypto/bcrypt"
)

func EncriptarPassword(pass string) (string, error) {

	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
