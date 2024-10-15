# Utiliza una imagen oficial de Go como base
FROM golang:1.21-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto de los archivos del proyecto al contenedor
COPY . .

# Compila el cliente
RUN go build -o client ./client.go

# Ejecuta el cliente cuando el contenedor se inicie
CMD ["./client"]
