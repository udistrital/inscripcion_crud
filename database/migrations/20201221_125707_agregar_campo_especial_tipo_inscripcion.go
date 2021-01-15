package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCampoEspecialTipoInscripcion_20201221_125707 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCampoEspecialTipoInscripcion_20201221_125707{}
	m.Created = "20201221_125707"

	migration.Register("AgregarCampoEspecialTipoInscripcion_20201221_125707", m)
}

// Run the migrations
func (m *AgregarCampoEspecialTipoInscripcion_20201221_125707) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201221_125707_agregar_campo_especial_tipo_inscripcion.up.sql")
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
func (m *AgregarCampoEspecialTipoInscripcion_20201221_125707) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201221_125707_agregar_campo_especial_tipo_inscripcion.down.sql")

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
