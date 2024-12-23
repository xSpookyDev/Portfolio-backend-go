package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

var Db *sql.DB

// Conectar abre una conexión con la base de datos MySQL.
func Conectar() {
	// Configuración de la base de datos directamente en el código
	dbUser := "demouser"
	dbPassword := "Dalebulla7"
	dbServer := "mysqlflex7.mysql.database.azure.com"
	dbPort := "3306"
	dbName := "golang-message"

	// Crear la cadena de conexión MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbServer, dbPort, dbName)

	// Conectar a la base de datos
	conection, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Error al conectar con la base de datos")
	}

	Db = conection
}

// CerrarConexion cierra la conexión actualmente abierta con la base de datos.
func CerrarConexion() {
	if err := Db.Close(); err != nil {
		log.Error().Err(err).Msg("Error al cerrar la conexión a la base de datos")
	}
}
