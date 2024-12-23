# Usa una imagen base ligera como Alpine
FROM alpine:latest

# Instala ca-certificates para HTTPS (si es necesario)
RUN apk --no-cache add ca-certificates

# Copia el archivo ejecutable desde tu máquina local al contenedor
COPY main.exe C:\Users\Asku\Desktop\proyecto3\main.exe

# Expone el puerto en el que la app estará corriendo
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./usr/local/bin/main.exe"]
