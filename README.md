# LoginEcho
Al ejecutar la app, hay que entrar al http://url:3000 que es la pagina de login
un se puede probar con el usur redis y pass: 12345678, realiza una consulta a una base externa en mongoatlas el cual una vez autenticado 
devuelve un token.
para saber los senadores de congreso de los EEUU 
se ingresa con un GET http//:127.0.0.1:3000/senadores y autenticando con el token que nos entrego despues del login
tambien se puede cononocer a los representantes de la camara del congresso tambien realizando un GET
a http://127.0.0.1:3000/house  en ambios casos devuelve un json

mas adelante quiero que al loguear por medio de una sesion se pueda mostrar de esta manera con algunas mejoras en el html
https://postimg.cc/bsZCLT1B
