package main

import (
	"errors"
	"fmt"
)

type Nodo struct {
	Titulo   string
	Director string
	Year     int
	Sig      *Nodo
}

type Lista struct {
	head *Nodo
}

type Cola struct {
	head *Nodo
	tail *Nodo
}

func (l *Lista) AgregarNodo(titulo string, director string, year int) {
	var nuevoNodo *Nodo = &Nodo{
		Titulo:   titulo,
		Director: director,
		Year:     year,
		Sig:      l.head,
	}

	l.head = nuevoNodo
}

func (l *Lista) Mostrar() {
	var actual *Nodo = l.head
	var i int = 0
	for actual != nil {
		fmt.Printf("ID: %d\nTitulo de la pelicula: %s\nDirector: %s\nAño: %d\n", i, actual.Titulo, actual.Director, actual.Year)
		for i := 0; i < len(actual.Titulo)+23; i++ {
			fmt.Print("-")
		}
		fmt.Println("\n\n")
		actual = actual.Sig
		i++
	}
	fmt.Println(l.head)
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

	fmt.Println(l.head)
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

// Cola
func (c *Cola) Encolar(titulo string, director string, year int) {
	nuevoNodo := &Nodo{
		Titulo:   titulo,
		Director: director,
		Year:     year,
		Sig:      nil,
	}

	//Si es el final de la cola
	if c.head == nil {
		c.head = nuevoNodo
		return
	}

	actual := c.head
	for ; actual.Sig != nil; actual = actual.Sig {
	}
	actual.Sig = nuevoNodo
}

func (c *Cola) Desencolar() {
	c.head = c.head.Sig

	if c.head == nil {
		c.tail = nil
	}
}

func (c *Cola) Mostrar() {
	var actual *Nodo = c.head
	var i int = 0
	for actual != nil {
		fmt.Printf("ID: %d\nTitulo de la pelicula: %s\nDirector: %s\nAño: %d\n", i, actual.Titulo, actual.Director, actual.Year)
		for i := 0; i < len(actual.Titulo)+23; i++ {
			fmt.Print("-")
		}
		fmt.Println("\n\n")
		actual = actual.Sig
		i++
	}
}
