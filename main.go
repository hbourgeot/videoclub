package main

import (
	"database/sql"
	"fmt"
	"jose/project/internal/models"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	var usuario, clave string
	var err error
	var bd *sql.DB
	var VideoClub *Lista = &Lista{}
	var VideoClubAlquiladas *Cola = &Cola{}

	bd, err = abrirBD()
	if err != nil {
		log.Println(err)
		return
	}

	for {
		var opc int
		fmt.Println("****** INICIAR SESIÓN ******")
		fmt.Println("\n\t1. Log in")
		fmt.Println("\t2. Salir\n")
		fmt.Print("Ingresa una opción: ")
		_, err = fmt.Scanln(&opc)
		if err != nil {
			log.Println(err)
			continue
		}

		switch opc {
		case 1:
			fmt.Print("Ingrese el usuario: ")
			_, err = fmt.Scanln(&usuario)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Print("Ingrese la clave: ")
			_, err = fmt.Scanln(&clave)
			if err != nil {
				log.Println(err)
				continue
			}

			var user *models.VideoClubUserModel = &models.VideoClubUserModel{BD: bd}

			if user.Authenticate(usuario, clave) {
				var nivel int = user.ObtenerNivel(usuario)
				switch nivel {
				case 1:
					standard(VideoClub, VideoClubAlquiladas)
					break
				case 2:
					admin(VideoClub)
					break
				}
			} else {
				fmt.Println("Credenciales incorrectas, adiós")
				continue
			}

		case 2:
			fmt.Print("\nBye")
			return
		}
	}
}

func abrirBD() (*sql.DB, error) {
	var bd *sql.DB
	var err error

	bd, err = sql.Open("mysql", "root:Wini.h16b.@/videoclub")
	if err != nil {
		return nil, err
	}

	if err = bd.Ping(); err != nil {
		return nil, err
	}

	return bd, nil
}
