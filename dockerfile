# Etapa de construcción
FROM golang:1.23.2 AS builder

# Establece el directorio de trabajo en el contenedor
WORKDIR /app

# Copia el archivo go.mod y go.sum para aprovechar la cache de Docker
COPY go.mod go.sum ./

# Descarga las dependencias
RUN go mod tidy

# Copia el código fuente de la aplicación al contenedor
COPY . .

# Copia las plantillas al contenedor
COPY templates /app/templates

# Compila la aplicación como binario estático
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Imagen final basada en Debian
FROM debian:bullseye-slim

# Copia el binario desde la etapa de construcción
COPY --from=builder /app/main /app/main

# Establece el directorio de trabajo
WORKDIR /app

# Asegura permisos de ejecución del binario
RUN chmod +x /app/main

# Expone el puerto en el que la aplicación escuchará
EXPOSE 8080

# Comando por defecto para ejecutar la aplicación
CMD ["/app/main"]
