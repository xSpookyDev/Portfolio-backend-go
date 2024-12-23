package main

import (
	"cmp"
	"fmt"
	"net/http"
	"os"
	v1 "proyecto_3/api/v1"
	"proyecto_3/connection"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339, NoColor: false})

	r := mux.NewRouter()

	r.HandleFunc("/", v1.Home).Methods("GET", "POST")
	r.HandleFunc("/listar", v1.Mysql_listar).Methods("GET", "POST")
	r.HandleFunc("/crear", v1.Mysql_crear).Methods("GET")
	r.HandleFunc("/crear_post", v1.Mysql_crear_post).Methods("POST")
	r.HandleFunc("/editar/{id:.*}", v1.Mysql_editar).Methods("GET")
	r.HandleFunc("/editar_post/{id:.*}", v1.Mysql_editar_post).Methods("POST")
	r.HandleFunc("/eliminar/{id:.*}", v1.Mysql_eliminar)

	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	r.PathPrefix("/public/").Handler(s)

	connection.Conectar()

	port := cmp.Or(os.Getenv("PORT"), "3000")

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port), // Escuchar en todas las interfaces
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info().Msg("Server running on port :" + port)
	server.ListenAndServe()
}
