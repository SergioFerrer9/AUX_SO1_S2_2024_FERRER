# Imagen base de Go
FROM golang:1.24.1

# Establecer el directorio
WORKDIR /app

# Copiar los archivos del cliente al contenedor
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar el cliente
RUN go build -o grpc-client client.go

# Ejecutar el cliente
CMD ["/app/grpc-client"]
