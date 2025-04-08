# Imagen base de Go
FROM golang:1.24.1

# Establecer el directorio
WORKDIR /app

# Copiar los archivos del servidor al contenedor
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar el servidor
RUN go build -o grpc-server server.go

# Exponer el puerto gRPC
EXPOSE 50051

# Ejecutar el servidor
CMD ["/app/grpc-server"]
