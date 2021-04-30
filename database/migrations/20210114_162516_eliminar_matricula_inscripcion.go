package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EliminarMatriculaInscripcion_20210114_162516 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EliminarMatriculaInscripcion_20210114_162516{}
	m.Created = "20210114_162516"

	migration.Register("EliminarMatriculaInscripcion_20210114_162516", m)
}

// Run the migrations
func (m *EliminarMatriculaInscripcion_20210114_162516) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210114_162516_eliminar_matricula_inscripcion.up.sql")
	if err != nil {
		//handle error
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
func (m *EliminarMatriculaInscripcion_20210114_162516) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210114_162516_eliminar_matricula_inscripcion.down.sql")
	if err != nil {
		//handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
