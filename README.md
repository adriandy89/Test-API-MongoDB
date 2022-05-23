# API en GO con base de datos mongodb

### Pasos:
1- Instalar Go

(Si está en algun país bloqueado, utilizar VPN para descargar e instalar dependencias necesarias)

2- Abrir una consola e instalar las dependencias (aparecen en el archivo 'go.mod'): 
### Dependencias necesarias

- go get github.com/cavaliergopher/grab/v3

- go get github.com/dgrijalva/jwt-go 

- go get github.com/go-ldap/ldap/v3

- go get github.com/gorilla/mux

- go get github.com/joho/godotenv

- go get github.com/rs/cors

- go get go.mongodb.org/mongo-driver

- go get golang.org/x/crypto


3- Correr el proyecto:
- go run main.go

4- Crear ejecutable (Opcional, y se puede configuar diferentes opciones para compilar para Linux, Windows o Mac):
- go build