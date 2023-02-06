package main

import (
	"fmt"
	"log"
	"time"
)

func admin(VideoClub *Lista) {
	var opc int
	var err error
	for {
		fmt.Println("\t******************************")
		fmt.Println("\t******************************")
		fmt.Println("\t***** Video Club Marrero *****")
		fmt.Println("\tBienvenido, administrador.")

		time.Sleep(3 * time.Second)

		fmt.Println("\n1. Agregar película")
		fmt.Println("2. Quitar película")
		fmt.Println("3. Actualizar película")
		fmt.Println("4. Mostrar todas las películas")
		fmt.Println("5. Cerrar sesión")
		fmt.Print("\nIngrese una opción: ")
		_, err = fmt.Scanln(&opc)
		if err != nil {
			log.Println(err)
			continue
		}

		switch opc {
		case 1:
			var titulo, director string
			var year int
			fmt.Print("\nIngrese el titulo de la pelicula: ")
			_, err = fmt.Scanln(&titulo)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Print("\nIngrese el director de la pelicula: ")
			_, err = fmt.Scanln(&director)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Print("\nIngrese el año de la pelicula: ")
			_, err = fmt.Scanln(&year)
			if err != nil {
				log.Println(err)
				continue
			}

			VideoClub.AgregarNodo(titulo, director, year)
			break
		case 2:
			var titulo string
			fmt.Print("\nIngrese el titulo de la pelicula: ")
			_, err = fmt.Scanln(&titulo)
			if err != nil {
				log.Println(err)
				continue
			}

			VideoClub.QuitarNodo(titulo)
			break
		case 3:
			var titulo, nuevoTitulo, nuevoDirector string
			var nuevoYear int
			fmt.Print("\nIngrese el titulo de la pelicula: ")
			_, err = fmt.Scanln(&titulo)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Print("\nIngrese el nuevo titulo de la pelicula: ")
			_, err = fmt.Scanln(&nuevoTitulo)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Print("\nIngrese el nuevo director de la pelicula: ")
			_, err = fmt.Scanln(&nuevoDirector)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Print("\nIngrese el nuevo año de la pelicula: ")
			_, err = fmt.Scanln(&nuevoYear)
			if err != nil {
				log.Println(err)
				continue
			}

			err = VideoClub.ActualizarNodo(titulo, nuevoTitulo, nuevoDirector, nuevoYear)
			if err != nil {
				log.Println(err)
				continue
			}
			break
		case 4:
			VideoClub.Mostrar()
			time.Sleep(5 * time.Second)
			break
		case 5:
			return
		default:
			fmt.Println("Opción inválida")
			break
		}

	}
}

func standard(VideoClub *Lista, VideoClubAlquilada *Cola) {
	var opc int
	var titulo string
	var err error

	for {
		fmt.Println("\t******************************")
		fmt.Println("\t***** Video Club Marrero *****")

		time.Sleep(3 * time.Second)

		fmt.Println("\n1. Buscar película")
		fmt.Println("2. Alquilar una película")
		fmt.Println("3. Regresar una película")
		fmt.Println("4. Mostrar peliculas alquiladas")
		fmt.Println("5. Cerrar sesión")
		fmt.Print("\nIngrese una opción: ")
		_, err = fmt.Scanln(&opc)
		if err != nil {
			log.Println(err)
			continue
		}
		switch opc {
		case 1:
			fmt.Print("Ingrese titulo de la pelicula: ")
			_, err = fmt.Scanln(&titulo)
			if err != nil {
				log.Println(err)
				continue
			}

			var pelicula *Nodo
			fmt.Println(titulo)
			pelicula, err = VideoClub.BuscarNodo(titulo)

			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Printf("\nTitulo: %s\nDirector: %s\nAño: %d\n", pelicula.Titulo, pelicula.Director, pelicula.Year)
			break
		case 2:
			fmt.Print("Ingrese titulo de la pelicula: ")
			_, err = fmt.Scanln(&titulo)
			if err != nil {
				log.Println(err)
				continue
			}

			var pelicula *Nodo
			pelicula, err = VideoClub.BuscarNodo(titulo)

			if err != nil {
				log.Println(err)
				continue
			}

			VideoClubAlquilada.Encolar(pelicula.Titulo, pelicula.Director, pelicula.Year)
			break
		case 3:
			VideoClubAlquilada.Desencolar()
			break
		case 4:
			VideoClubAlquilada.Mostrar()
			time.Sleep(5 * time.Second)
			break
		case 5:
			return
		}
	}
}
