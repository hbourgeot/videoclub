package main

import (
	"database/sql"
	"fmt"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	var usuario, clave string
	var err error

	for {
		var opc int
		fmt.Println("****** INICIAR SESIÓN ******")
		fmt.Println("\t\n1. Log in")
		fmt.Println("\t\n2. Salir\n")
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

			//var bd *sql.DB
			_, err = abrirBD(usuario, clave)
			if err != nil {
				log.Println(err)
				continue
			}
		case 2:
			fmt.Print("\nBye")
			return
		}
	}
}

func abrirBD(usuario, clave string) (*sql.DB, error) {
	var bd *sql.DB
	var err error

	bd, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/videoClub", usuario, clave))
	if err != nil {
		return nil, err
	}

	if err = bd.Ping(); err != nil {
		return nil, err
	}

	return bd, nil
}
