package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/soUNds-unal/sounds/sounds_ms/sounds_auth_ms/models"
)

func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("Estaesunapruebadesguridad")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"ubicacion":        t.Ubicacion,
		"sitio_web":        t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
