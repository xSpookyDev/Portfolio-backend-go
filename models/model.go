package models

type Usuario struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Contrasena string `json:"contrasena"`
	Comentario string `json:"comentario"`
}

type Usuarios []Usuario
