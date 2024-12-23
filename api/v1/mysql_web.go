package v1

import (
	"fmt"
	"html/template"
	"net/http"
	"proyecto_3/connection"
	"proyecto_3/models"

	"github.com/gorilla/mux"
)

func Mysql_listar(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/newMessage.html"))

	connection.Conectar()
	sql := "Select id, nombre, apellido, comentario from usuarios order by id desc"
	clientes := models.Usuarios{}
	datos, err := connection.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer connection.CerrarConexion()
	for datos.Next() {
		dato := models.Usuario{}
		datos.Scan(&dato.ID, &dato.Nombre, &dato.Apellido, &dato.Comentario)
		clientes = append(clientes, dato)
	}

	data := map[string]interface{}{
		"Datos": clientes,
	}

	tmpl.Execute(w, data)
}

func Mysql_crear(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/crear.html"))

	connection.Conectar()
	sql := "insert into usuarios values (null, ?, ?, ?;)"
	_, err := connection.Db.Exec(sql, r.FormValue("name"), r.FormValue("apellido"), r.FormValue("contrasena"), r.FormValue("comentario"))
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, nil)
}

func Mysql_crear_post(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	connection.Conectar()
	defer connection.CerrarConexion()

	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	contrasena := r.FormValue("contrasena")
	comentario := r.FormValue("comentario")

	if nombre == "" || apellido == "" || contrasena == "" {
		http.Error(w, "Faltan campos requeridos", http.StatusBadRequest)
		return
	}

	fmt.Println("Recibidos: Nombre:", nombre, "Apellido:", apellido, "Contraseña:", contrasena, "Comentario:", comentario)

	sql := "INSERT INTO usuarios (nombre, apellido, contrasena, comentario) VALUES (?, ?, ?, ?)"

	_, err := connection.Db.Exec(sql, nombre, apellido, contrasena, comentario)
	if err != nil {
		http.Error(w, "Error al insertar datos: "+err.Error(), http.StatusInternalServerError)
		fmt.Println("Error en la inserción:", err)
		return
	}

	http.Redirect(w, r, "/#testimonials", http.StatusSeeOther)
}

func Mysql_editar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "ID no proporcionado", http.StatusBadRequest)
		return
	}

	connection.Conectar()
	defer connection.CerrarConexion()

	sql := "SELECT id, nombre, apellido, contrasena, comentario FROM usuarios WHERE id = ?"
	var usuario models.Usuario
	err := connection.Db.QueryRow(sql, id).Scan(&usuario.ID, &usuario.Nombre, &usuario.Apellido, &usuario.Contrasena, &usuario.Comentario)
	if err != nil {
		http.Error(w, "Error al obtener datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/editar.html"))
	tmpl.Execute(w, usuario)
}

func Mysql_editar_post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")
	contrasena := r.FormValue("contrasena")
	comentario := r.FormValue("comentario")

	if nombre == "" || apellido == "" || contrasena == "" {
		http.Error(w, "Faltan campos requeridos", http.StatusBadRequest)
		return
	}

	connection.Conectar()
	defer connection.CerrarConexion()

	sql := "UPDATE usuarios SET nombre = ?, apellido = ?, contrasena = ?, comentario = ? WHERE id = ?"

	_, err := connection.Db.Exec(sql, nombre, apellido, contrasena, comentario, id)
	if err != nil {
		http.Error(w, "Error al actualizar datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/listar", http.StatusSeeOther)
}

func Mysql_eliminar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	connection.Conectar()
	defer connection.CerrarConexion()

	sql := "DELETE FROM usuarios WHERE id = ?"
	_, err := connection.Db.Exec(sql, id)
	if err != nil {
		http.Error(w, "Error al eliminar el registro: "+err.Error(), http.StatusInternalServerError)
		fmt.Println("Error al eliminar el registro:", err)
		return
	}

	http.Redirect(w, r, "/listar", http.StatusSeeOther)
}
