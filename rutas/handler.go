package rutas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	"github.com/barbaselperro/Echo/base"
	"github.com/barbaselperro/Echo/cacheRedis"
	"github.com/barbaselperro/Echo/tokens"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// leemos el archivo donde esta escrita la clave para la firma del token en este caso clave.txt
var file, _ = ioutil.ReadFile("clave.txt")
var key = []byte(file)

func Index(c echo.Context) error {
	html, _ := template.ParseFiles("login.html")
	nombre := "Golang"
	fmt.Println(file)
	return html.Execute(c.Response(), nombre)
}

func Tutoken(c echo.Context) error {
	c.Request().ParseForm()
	usuario := c.FormValue("username")
	clave := c.FormValue("password")
	datomongo := base.ObtieneMongo(usuario)

	if usuario == datomongo.User && clave == datomongo.Pass {
		mytoken := tokens.ReciboToken(key, datomongo)
		var tuToken = make(map[string]string)
		tuToken["token"] = mytoken
		fmt.Println(tuToken)
		return c.JSON(http.StatusOK, tuToken)
	}
	return c.String(http.StatusOK, "Usuario o clave incorrecta")
}

func Acceso(c echo.Context) error {
	token := strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")[1]
	fmt.Println("el token que se recibe es: ", token)
	tokenOk, _ := jwt.Parse(token, func(tokens *jwt.Token) (interface{}, error) {
		return key, nil
	})
	fmt.Println(tokenOk.Claims)
	fmt.Println(tokenOk.Valid)
	if tokenOk.Valid {
		return c.String(http.StatusOK, "Ok")
	}
	return c.String(http.StatusOK, "no tiene acceso")
}

func Tenadores(c echo.Context) error {
	jsn := cacheRedis.Consulta("senadores")
	if jsn == "" {
		datos := base.Senadores()
		cacheRedis.RedisDB(datos)
		return headerjwt(c, datos)
	} else {
		var jsnStr []base.Senates
		json.Unmarshal([]byte(jsn), &jsnStr)
		return headerjwt(c, jsnStr)
	}

}

func House(c echo.Context) error {
	datos := base.House()
	return headerjwt(c, datos)
}

func headerjwt(c echo.Context, datos interface{}) error {
	var token string
	//en este ciclo se verifica si se recibe el token en el header
	if c.Request().Header.Get("Authorization") != "" {
		token = strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")[1]
	} else {
		return c.String(http.StatusOK, "no esta autorizado para ingresar sin token")
	}

	//en esta seccion se realiza el parse del token mejor dicho es para verificar
	//si el token esta correcto
	tokenOk, _ := jwt.Parse(token, func(tokens *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if tokenOk.Valid {
		//datos := base.Senadores()
		//fmt.Println(datos)
		return json.NewEncoder(c.Response()).Encode(datos)
	} else {
		return c.String(http.StatusOK, "no esta autorizado para poder ingresar")
	}
}
