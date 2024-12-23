package connection

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var Db *sql.DB

// Conectar abre una conexion con la base de datos MySQL.
func Conectar() {

	errorVariables := godotenv.Load()
	if errorVariables != nil {
		log.Fatal().Err(errorVariables).Msg("Error al cargar el archivo .env")
	}

	conection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
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
