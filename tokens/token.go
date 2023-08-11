package tokens

import (
	"fmt"
	"time"

	"github.com/barbaselperro/Echo/base"
	"github.com/golang-jwt/jwt/v5"
)

func ReciboToken(key []byte, dato base.Usuario) string {
	//Primero creamos una variable que es un struct de jwt.Token
	var t *jwt.Token

	//Con NewTithClaims el primer argunmento es el tipo de algoritmo para la firma
	//el segundo parametro es el payload o el contenido mejor conocido como los claims que son los datos
	//del usuario por ejemplo nombre de user tiempo de expiracion y otros datos adicionales
	var claims jwt.RegisteredClaims = jwt.RegisteredClaims{Issuer: dato.User, ID: dato.Id, ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute))}
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //jwt.MapClaims{"user": "barbas", "password": "Barbaselperro"})

	//el metedo SignedString sirve para firmar nuestro token con una clave, el cual nos devuelve un token
	//que se divede en 3 partes XXXXX.YYYYYY.ZZZZZZZ El header contine el tipo de algoritmo, despues el
	//contenido o claims informacion del usuario y por ultimo el verify signature o firma digital que se
	//obtiene por medio de las dos secciones anteriores y permite analizar si el contenido fue modificado. El cual
	//fue firmado usando una clave y el metodos SignedString en este caso se uso asi:
	// t.SignetString(key) es una variable que contiene una clave creada
	str, _ := t.SignedString(key)
	return str
}

func Verificar(token string, key byte) int {
	tokenOk, err := jwt.Parse(token, func(tokens *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if tokenOk.Valid {
		return 1
	} else {
		fmt.Println(err)
		return 0
	}

}
