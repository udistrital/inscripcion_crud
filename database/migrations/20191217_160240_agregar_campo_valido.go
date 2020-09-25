package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCampoValido_20191217_160240 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCampoValido_20191217_160240{}
	m.Created = "20191217_160240"

	migration.Register("AgregarCampoValido_20191217_160240", m)
}

// Run the migrations
func (m *AgregarCampoValido_20191217_160240) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20191217_160240_agregar_campo_valido.up.sql")

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
func (m *AgregarCampoValido_20191217_160240) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20191217_160240_agregar_campo_valido.down.sql")

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
