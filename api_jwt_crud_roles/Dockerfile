# Imagen base
FROM golang:1.24.5

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar módulos de Go y descargar dependencias primero (cache)
COPY go.mod go.sum ./
RUN go mod tidy

# Copiar todo el proyecto
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer puerto de la API
EXPOSE 8080

# Comando para ejecutar la API
CMD ["./main"]
