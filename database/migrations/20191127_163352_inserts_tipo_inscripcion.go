package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertsTipoInscripcion_20191127_163352 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsTipoInscripcion_20191127_163352{}
	m.Created = "20191127_163352"

	migration.Register("InsertsTipoInscripcion_20191127_163352", m)
}

// Run the migrations
func (m *InsertsTipoInscripcion_20191127_163352) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20191127_163352_inserts_tipo_inscripcion.up.sql")

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
func (m *InsertsTipoInscripcion_20191127_163352) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20191127_163352_inserts_tipo_inscripcion.down.sql")

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
