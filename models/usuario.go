package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" "json:id"`
	Nombre          string             `bson:"nombre" json:"nombre, omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos, omitempty"`
	FechaNacimiento time.Time          `bson:"fecNac" json:"fecNac, omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"pass" json:"pass, ometempty"`
	Avatar          string             `bson:"avatar" json:"avatar, omitempty"`
	banner          string             `bson:"banner"	json:"banner", omitempty"`
}
