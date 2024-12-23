package v1

import (
	"fmt"
	"html/template"
	"net/http"
	"proyecto_3/connection"
	"proyecto_3/models"
	"proyecto_3/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		utils.SendEmail(name, email, message)
	}

	connection.Conectar()
	sql := "SELECT id, nombre, apellido, comentario FROM usuarios ORDER BY id DESC"
	clientes := models.Usuarios{}
	datos, err := connection.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error al obtener los usuarios", http.StatusInternalServerError)
		return
	}
	defer connection.CerrarConexion()

	for datos.Next() {
		dato := models.Usuario{}
		err := datos.Scan(&dato.ID, &dato.Nombre, &dato.Apellido, &dato.Comentario)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error al leer los datos", http.StatusInternalServerError)
			return
		}
		clientes = append(clientes, dato)
	}

	data := map[string]interface{}{
		"Usuarios": clientes,
	}

	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, data)
}
