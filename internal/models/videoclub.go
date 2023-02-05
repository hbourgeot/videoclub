package models

import (
	"database/sql"
	"log"
)

type VideoClubUser struct {
	ID       int
	Username string
	Clave    string
	Nivel    string
}

type VideoClubUserModel struct {
	BD *sql.DB
}

func (m *VideoClubUserModel) Authenticate(username, clave string) bool {
	var consulta string = "SELECT * FROM usuario WHERE username = ? AND contraseña = ?"
	var registro *sql.Row = m.BD.QueryRow(consulta, username, clave)
	var err error
	var v *VideoClubUser = &VideoClubUser{}
	if err = registro.Scan(&v.ID, &v.Username, &v.Clave, &v.Nivel); err != nil {
		return false
	}
	return true
}

func (m *VideoClubUserModel) ObtenerNivel(username string) int {
	var consulta string = "SELECT nivel FROM usuario WHERE username = ?"
	var registro *sql.Row = m.BD.QueryRow(consulta, username)
	var err error
	var v *VideoClubUser = &VideoClubUser{}
	if err = registro.Scan(&v.Nivel); err != nil {
		log.Println(err)
		return 0
	}

	var rol int

	if v.Nivel == "admin" {
		rol = 2
	} else {
		rol = 1
	}

	return rol
}
