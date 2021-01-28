package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertEstadoInscripcion_20210128_071819 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertEstadoInscripcion_20210128_071819{}
	m.Created = "20210128_071819"

	migration.Register("InsertEstadoInscripcion_20210128_071819", m)
}

// Run the migrations
func (m *InsertEstadoInscripcion_20210128_071819) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210128_071819_insert_estado_inscripcion.up.sql")

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
func (m *InsertEstadoInscripcion_20210128_071819) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210128_071819_insert_estado_inscripcion.down.sql")

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
