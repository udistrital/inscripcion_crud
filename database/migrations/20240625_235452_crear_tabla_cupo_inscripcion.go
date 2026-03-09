package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaCupoInscripcion_20240625_235452 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaCupoInscripcion_20240625_235452{}
	m.Created = "20240625_235452"

	migration.Register("CrearTablaCupoInscripcion_20240625_235452", m)
}

// Run the migrations
func (m *CrearTablaCupoInscripcion_20240625_235452) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20240625_235452_crear_tabla_cupo_inscripcion.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *CrearTablaCupoInscripcion_20240625_235452) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20240625_235452_crear_tabla_cupo_inscripcion.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
