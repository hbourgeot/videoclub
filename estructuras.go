package main

import (
	"errors"
	"fmt"
)

type Nodo struct {
	ID       int
	Titulo   string
	Director string
	Year     int
	Sig      *Nodo
}

type Lista struct {
	head *Nodo
}

func (l *Lista) AgregarNodo(id int, titulo string, director string, year int) {
	var nuevoNodo *Nodo = &Nodo{
		Titulo:   titulo,
		Director: director,
		Year:     year,
		ID:       id,
		Sig:      l.head,
	}

	l.head = nuevoNodo
	fmt.Printf("Creado nuevo Nodo con el ID #%d", nuevoNodo.ID)
}

func (l *Lista) Mostrar() {
	var actual *Nodo = l.head
	for actual != nil {
		fmt.Printf("ID: %d\nTitulo de la pelicula: %s\nDirector: %s\nAño: %d\n")
		for i := 0; i < len(actual.Titulo)+23; i++ {
			fmt.Print("-")
		}
		fmt.Println("\n\n")
	}
}

func (l *Lista) QuitarNodo(titulo string) {
	var actual, previo *Nodo = l.head, nil

	for actual != nil {
		if actual.Titulo == titulo {
			if previo == nil {
				l.head = actual.Sig
			} else {
				previo.Sig = actual.Sig
			}

			actual = nil // se borra de esta forma el Nodo actual
			break
		}

		previo = actual
		actual = actual.Sig
	}
}

func (l *Lista) BuscarNodo(titulo string) (*Nodo, error) {
	var temp *Nodo = l.head
	for temp != nil {
		if temp.Titulo == titulo {
			return temp, nil
		}

		temp = temp.Sig
	}

	var err error = errors.New("Pelicula no encontrada")
	return nil, err
}

func (l *Lista) ActualizarNodo(titulo, nuevoTitulo, nuevoDirector string, nuevoYear int) error {
	var actual *Nodo
	var err error
	actual, err = l.BuscarNodo(titulo)
	if err != nil {
		return err
	}

	actual.Titulo = nuevoTitulo
	actual.Director = nuevoDirector
	actual.Year = nuevoYear
	return nil
}
