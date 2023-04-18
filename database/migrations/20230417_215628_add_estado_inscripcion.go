package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddEstadoInscripcion_20230417_215628 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddEstadoInscripcion_20230417_215628{}
	m.Created = "20230417_215628"

	migration.Register("AddEstadoInscripcion_20230417_215628", m)
}

// Run the migrations
func (m *AddEstadoInscripcion_20230417_215628) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20230417_215628_add_estado_inscripcion.up.sql")

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
func (m *AddEstadoInscripcion_20230417_215628) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20230417_215628_add_estado_inscripcion.down.sql")

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
