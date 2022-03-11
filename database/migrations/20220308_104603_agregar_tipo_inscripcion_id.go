package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarTipoInscripcionId_20220308_104603 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarTipoInscripcionId_20220308_104603{}
	m.Created = "20220308_104603"

	migration.Register("AgregarTipoInscripcionId_20220308_104603", m)
}

// Run the migrations
func (m *AgregarTipoInscripcionId_20220308_104603) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20220308_104603_agregar_tipo_inscripcion_id.up.sql")

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
func (m *AgregarTipoInscripcionId_20220308_104603) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20220308_104603_agregar_tipo_inscripcion_id.down.sql")

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
